# Error Handling Guide

Learn how to effectively use GoSugar's panic-based error handling system.

## 📋 Contents

- [General Approach](#general-approach)
- [Panic vs Error](#panic-vs-error)
- [Strategies](#strategies)
- [Real-World Examples](#real-world-examples)

---

## General Approach

### GoSugar's Philosophy

GoSugar's error handling is based on this principle:

```
Critical errors → Panic (early notification)
Optional operations → Try/Or (graceful fallback)
```

### Two-Level Strategy

```
┌─────────────────────────────────┐
│  Startup Phase (Initialization) │
│  ← Panic appropriate            │
└─────────────────────────────────┘
              ↓
    - Config validation
    - Required env vars check
    - Setup errors

┌─────────────────────────────────┐
│  Runtime Phase (Execution)      │
│  ← Try/Or more appropriate      │
└─────────────────────────────────┘
              ↓
    - Optional file operations
    - API calls
    - Network operations
```

---

## Panic vs Error

### When to Panic?

✅ **Use Panic:**

```go
// 1. Configuration error
apiKey := gosugar.MustEnv("API_KEY")  // Panics if missing

// 2. Startup validation
port := gosugar.EnvInt("PORT", 0)
if port == 0 {
    panic("PORT must be > 0")
}

// 3. Input validation (CLI)
username := gosugar.Input(
    "Username: ",
    gosugar.NotEmpty(),  // Panics on error
)
```

**Rationale:**
- These errors should not be silently ignored
- User or admin should know immediately
- Application shouldn't run in wrong state

### When to Use Try/Or?

✅ **Use Try/Or:**

```go
// 1. Optional files
content, ok := gosugar.Try(func() string {
    return gosugar.ReadFile("optional.json")
})
result := gosugar.Or(content, ok, "default")

// 2. Network operations
data, ok := gosugar.Try(func() string {
    return gosugar.MustGetBody("https://example.com")
})
if !ok {
    fmt.Println("Network error, using cache")
    data = loadFromCache()
}

// 3. Non-critical operations
_, ok := gosugar.Try(func() {
    gosugar.AppendFile("debug.log", "info")
    return true
})
// Even if fails, we don't care
```

**Rationale:**
- These operations can fail and we can continue
- Graceful degradation
- Better user experience

---

## Strategies

### Strategy 1: Strict Startup

```go
package main

import (
    "fmt"
    "os"
    "github.com/coderianx/gosugar"
)

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Fprintf(os.Stderr, "Fatal: %v\n", r)
            os.Exit(1)
        }
    }()

    // Required configuration
    fmt.Println("Loading configuration...")

    dbURL := gosugar.MustEnv("DATABASE_URL")
    apiKey := gosugar.MustEnv("API_KEY")
    port := gosugar.EnvInt("PORT", 0)

    if port == 0 {
        panic("PORT env var not set or invalid")
    }

    fmt.Println("Configuration OK")
    fmt.Println("Starting application...")

    // Application logic
    runApp(port, dbURL, apiKey)
}

func runApp(port int, db, api string) {
    fmt.Printf("Running on port %d\n", port)
}
```

### Strategy 2: Flexible Startup

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func loadConfig() map[string]interface{} {
    config := make(map[string]interface{})

    // Critical: ENV
    config["api_key"] = gosugar.MustEnv("API_KEY")

    // Semi-critical: File
    content, ok := gosugar.Try(func() string {
        return gosugar.ReadFile("config.json")
    })
    if ok {
        config["settings"] = content
    } else {
        config["settings"] = "{}"
        fmt.Println("Warning: config.json not found, using defaults")
    }

    // Optional: Port
    config["port"] = gosugar.EnvInt("PORT", 8080)

    return config
}

func main() {
    config := loadConfig()
    fmt.Printf("Config loaded: %v\n", config)
}
```

### Strategy 3: Error Chaining

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func processFile(filename string) (string, error) {
    // Step 1: Read file
    content, ok := gosugar.Try(func() string {
        return gosugar.ReadFile(filename)
    })

    if !ok {
        return "", fmt.Errorf("cannot read file: %s", filename)
    }

    // Step 2: Process
    processed := processContent(content)

    // Step 3: Write
    _, ok = gosugar.Try(func() {
        gosugar.WriteFile(filename+".processed", processed)
        return true
    })

    if !ok {
        return "", fmt.Errorf("cannot write processed file")
    }

    return processed, nil
}

func processContent(s string) string {
    // Processing logic
    return "processed: " + s
}

func main() {
    result, err := processFile("data.txt")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Println("Success:", result)
}
```

---

## Real-World Examples

### Example 1: Web Scraper

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func scrapeWebpage(url string) string {
    fmt.Printf("Scraping: %s\n", url)

    // Try to fetch
    body, ok := gosugar.Try(func() string {
        return gosugar.MustGetBody(url)
    })

    if !ok {
        fmt.Println("Network error, using cached version")
        // Fallback
        cached, ok := gosugar.Try(func() string {
            return gosugar.ReadFile("cache/" + url + ".html")
        })
        body = gosugar.Or(cached, ok, "<html>Error</html>")
    }

    // Log the fetch
    gosugar.AppendFile(
        "scraper.log",
        fmt.Sprintf("Fetched %s - %d bytes\n", url, len(body)),
    )

    return body
}

func main() {
    page := scrapeWebpage("https://example.com")
    fmt.Println("Got page:", len(page), "bytes")
}
```

### Example 2: Data Pipeline

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func pipeline() {
    fmt.Println("Starting data pipeline...")

    // Step 1: Config
    gosugar.EnvFile(".env")
    output := gosugar.EnvString("OUTPUT_DIR", "./output")

    // Step 2: Input file
    fmt.Println("Reading input...")
    input, ok := gosugar.Try(func() string {
        return gosugar.ReadFile("input.csv")
    })

    if !ok {
        fmt.Println("No input file, generating test data")
        input = generateTestData()
    }

    // Step 3: Process
    fmt.Println("Processing...")
    processed := processData(input)

    // Step 4: Output
    fmt.Println("Writing output...")
    _, ok = gosugar.Try(func() {
        gosugar.WriteFile(output+"/result.csv", processed)
        return true
    })

    if !ok {
        fmt.Println("Warning: Could not write output file")
    }

    fmt.Println("Pipeline complete!")
}

func generateTestData() string {
    return "id,name,value\n1,test,100\n"
}

func processData(data string) string {
    return "PROCESSED: " + data
}

func main() {
    pipeline()
}
```

### Example 3: Interactive Form with Error Handling

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
    "strings"
)

// Custom validator
func isValidEmail() gosugar.Validator {
    return func(s string) error {
        if !strings.Contains(s, "@") {
            return fmt.Errorf("invalid email")
        }
        return nil
    }
}

func collectUserData() {
    fmt.Println("=== User Registration ===\n")

    // Username: strictly validated
    username := gosugar.Input(
        "Username (3-20): ",
        gosugar.NotEmpty(),
        gosugar.MinLen(3),
        gosugar.MaxLen(20),
    )

    // Email: strict
    email := gosugar.Input(
        "Email: ",
        gosugar.NotEmpty(),
        isValidEmail(),
    )

    // Age: with fallback
    age := gosugar.InputInt("Age (18+): ", 18)
    if age < 18 {
        fmt.Println("Error: Must be 18+")
        return
    }

    // Bio: optional
    bio, ok := gosugar.Try(func() string {
        return gosugar.Input(
            "Bio (optional): ",
            gosugar.MaxLen(500),
        )
    })
    bio = gosugar.Or(bio, ok, "")

    // Save to file
    data := fmt.Sprintf(
        "User: %s\nEmail: %s\nAge: %d\nBio: %s\n",
        username, email, age, bio,
    )

    _, ok = gosugar.Try(func() {
        gosugar.AppendFile("users.txt", data)
        return true
    })

    if !ok {
        fmt.Println("Warning: Could not save user data")
    } else {
        fmt.Println("User registered successfully!")
    }
}

func main() {
    collectUserData()
}
```

---

## Debugging Tips

### Tip 1: Panic Stacktrace

```go
// To see stacktrace
defer func() {
    if r := recover(); r != nil {
        fmt.Printf("Panic: %v\n", r)
        // Stacktrace automatically printed
    }
}()

// Risky code
val := gosugar.Must(someFunc())
```

### Tip 2: Debugging with Try

```go
value, ok := gosugar.Try(func() string {
    return gosugar.ReadFile("data.txt")
})

if !ok {
    fmt.Println("Operation failed - debugging info:")
    // Fallback and logging
    fmt.Println("Trying alternative...")
}
```

### Tip 3: Conditional Panic

```go
// Panic if condition fails
port := gosugar.EnvInt("PORT", 0)

if port < 1 || port > 65535 {
    panic(fmt.Sprintf("Invalid port: %d", port))
}
```

---

## Best Practices Summary

| Rule | Examples |
|------|----------|
| **Startup: Panic** | EnvFile, MustEnv, validation |
| **Runtime: Try/Or** | File ops, HTTP, network |
| **Validation: Panic** | Input validators, config |
| **Optional: Try/Or** | Cache, logging, fallback |
| **Log everything** | AppendFile error details |
| **Graceful fallback** | Defaults, cached values |

---

## Related Modules

- [`../api/errors.md`](../api/errors.md) - Errors API reference
- [`design-decisions.md`](../architecture/design-decisions.md) - Why panic?
