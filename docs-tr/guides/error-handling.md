# Hata Yönetimi Rehberi

GoSugar'ın panic-based error handling sistemini nasıl etkili kullanılacağını öğrenin.

## 📋 İçindekiler

- [Genel Yaklaşım](#genel-yaklaşım)
- [Panic vs Error](#panic-vs-error)
- [Stratejiler](#stratejiler)
- [Real-World Örnekler](#real-world-örnekler)

---

## Genel Yaklaşım

### GoSugar'ın Felsefesi

GoSugar'ın hata yönetimi şu ilkeye dayanır:

```
Kritik hatalar → Panic (erken bildirim)
Opsiyonel işlemler → Try/Or (graceful fallback)
```

### İki Düzeyli Strateji

```
┌─────────────────────────────────┐
│  Startup Phase (Initialization) │
│  ← Panic uygun                  │
└─────────────────────────────────┘
              ↓
    - Config validation
    - Required env vars check
    - Setup errors

┌─────────────────────────────────┐
│  Runtime Phase (Execution)      │
│  ← Try/Or daha uygun            │
└─────────────────────────────────┘
              ↓
    - Optional file operations
    - API calls
    - Network operations
```

---

## Panic vs Error

### Ne Zaman Panic?

✅ **Panic Kullanın:**

```go
// 1. Konfigürasyon hatası
apiKey := gosugar.MustEnv("API_KEY")  // Yoksa panic

// 2. Startup validation
port := gosugar.EnvInt("PORT", 0)
if port == 0 {
    panic("PORT must be > 0")
}

// 3. Input validation (CLI)
username := gosugar.Input(
    "Username: ",
    gosugar.NotEmpty(),  // Hata varsa panic
)
```

**Gerekçe:**
- Bu hataları "sessizce" geçmek danger
- Kullanıcı veya admin anında öğrenmeli
- Uygulama yanlış state'te çalışmasın

### Ne Zaman Try/Or?

✅ **Try/Or Kullanın:**

```go
// 1. Opsiyonel dosyalar
content, ok := gosugar.Try(func() string {
    return gosugar.ReadFile("optional.json")
})
result := gosugar.Or(content, ok, "default")

// 2. Network işlemleri
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
// Başarısız olsa da umursamıyoruz
```

**Gerekçe:**
- Bu işlemler başarısız olsa da devam edebilirsin
- Graceful degradation
- User experience daha iyi

---

## Stratejiler

### Strateji 1: Strict Startup

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

    // Zorunlu konfigürasyonlar
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

### Strateji 2: Flexible Startup

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func loadConfig() map[string]interface{} {
    config := make(map[string]interface{})

    // Kritik: ENV
    config["api_key"] = gosugar.MustEnv("API_KEY")

    // Semi-kritik: File
    content, ok := gosugar.Try(func() string {
        return gosugar.ReadFile("config.json")
    })
    if ok {
        config["settings"] = content
    } else {
        config["settings"] = "{}"
        fmt.Println("Warning: config.json not found, using defaults")
    }

    // Opsiyonel: Port
    config["port"] = gosugar.EnvInt("PORT", 8080)

    return config
}

func main() {
    config := loadConfig()
    fmt.Printf("Config loaded: %v\n", config)
}
```

### Strateji 3: Error Chaining

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func processFile(filename string) (string, error) {
    // Adım 1: Dosya oku
    content, ok := gosugar.Try(func() string {
        return gosugar.ReadFile(filename)
    })

    if !ok {
        return "", fmt.Errorf("cannot read file: %s", filename)
    }

    // Adım 2: İşle
    processed := processContent(content)

    // Adım 3: Yaz
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

## Real-World Örnekler

### Örnek 1: Web Scraper

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

### Örnek 2: Data Pipeline

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func pipeline() {
    fmt.Println("Starting data pipeline...")

    // Adım 1: Config
    gosugar.EnvFile(".env")
    output := gosugar.EnvString("OUTPUT_DIR", "./output")

    // Adım 2: Input dosya
    fmt.Println("Reading input...")
    input, ok := gosugar.Try(func() string {
        return gosugar.ReadFile("input.csv")
    })

    if !ok {
        fmt.Println("No input file, generating test data")
        input = generateTestData()
    }

    // Adım 3: Process
    fmt.Println("Processing...")
    processed := processData(input)

    // Adım 4: Output
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

### Örnek 3: Interactive Form dengan Error Handling

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
// Stacktrace'i görmek için
defer func() {
    if r := recover(); r != nil {
        fmt.Printf("Panic: %v\n", r)
        // Stack trace otomatik yazdırılır
    }
}()

// Riskli kod
val := gosugar.Must(someFunc())
```

### Tip 2: Try ile Debugging

```go
value, ok := gosugar.Try(func() string {
    return gosugar.ReadFile("data.txt")
})

if !ok {
    fmt.Println("Operation failed - debugging info:")
    // Fallback ve loglama
    fmt.Println("Trying alternative...")
}
```

### Tip 3: Conditional Panic

```go
// Eğer koşul başarısızsa panic
port := gosugar.EnvInt("PORT", 0)

if port < 1 || port > 65535 {
    panic(fmt.Sprintf("Invalid port: %d", port))
}
```

---

## Best Practices Özeti

| Kural | Örnekler |
|-------|----------|
| **Startup: Panic** | EnvFile, MustEnv, validation |
| **Runtime: Try/Or** | File ops, HTTP, network |
| **Validation: Panic** | Input validators, config |
| **Optional: Try/Or** | Cache, logging, fallback |
| **Log everything** | AppendFile hata detayları |
| **Graceful fallback** | Defaults, cached values |

---

## İlgili Modüller

- [`../api/errors.md`](../api/errors.md) - Errors API referansı
- [`design-decisions.md`](../architecture/design-decisions.md) - Neden panic?

