# Tasarım Desenleri - GoSugar İle Yaygın Patterns

GoSugar'ı farklı senaryolarda nasıl kullanılacağını gösteren practical patterns.

## 📋 İçindekiler

- [Configuration Pattern](#configuration-pattern)
- [Input Validation Pattern](#input-validation-pattern)
- [Error Recovery Pattern](#error-recovery-pattern)
- [Test Data Generation](#test-data-generation)
- [CLI Application Pattern](#cli-application-pattern)

---

## Configuration Pattern

### Pattern: Environment-Based Configuration

Uygulama başında ortamı yükle ve doğrula.

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

type Config struct {
    AppName       string
    Port          int
    Debug         bool
    DatabaseURL   string
}

func loadConfig() Config {
    // .env dosyasını yükle
    gosugar.EnvFile(".env")

    // Zorunlu variables
    dbURL := gosugar.MustEnv("DATABASE_URL")

    // Opsiyonel variables (defaults ile)
    cfg := Config{
        AppName:     gosugar.EnvString("APP_NAME", "MyApp"),
        Port:        gosugar.EnvInt("PORT", 8080),
        Debug:       gosugar.EnvBool("DEBUG", false),
        DatabaseURL: dbURL,
    }

    return cfg
}

func main() {
    cfg := loadConfig()
    fmt.Printf("Starting %s on port %d\n", cfg.AppName, cfg.Port)
}
```

**Best Practice:**
```go
// Doğrulama
if cfg.Port < 1 || cfg.Port > 65535 {
    panic("Invalid port")
}

// Logging
if cfg.Debug {
    fmt.Println("Debug mode enabled")
}
```

---

## Input Validation Pattern

### Pattern 1: Multi-Field Form

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

type User struct {
    Username string
    Email    string
    Age      int
}

func collectUserInput() User {
    fmt.Println("=== User Registration ===\n")

    // Username: 3-20 caracteres
    username := gosugar.Input(
        "Username (3-20 chars): ",
        gosugar.NotEmpty(),
        gosugar.MinLen(3),
        gosugar.MaxLen(20),
    )

    // Email: 5-100 caracteres
    email := gosugar.Input(
        "Email (5-100 chars): ",
        gosugar.NotEmpty(),
        gosugar.MinLen(5),
        gosugar.MaxLen(100),
    )

    // Age: integer
    age := gosugar.InputInt("Age (18+): ", 0)
    if age < 18 {
        panic("Must be 18 or older")
    }

    return User{
        Username: username,
        Email:    email,
        Age:      age,
    }
}

func main() {
    user := collectUserInput()
    fmt.Printf("\nUser: %s (%s), Age: %d\n", user.Username, user.Email, user.Age)
}
```

### Pattern 2: Custom Validators

```go
package main

import (
    "fmt"
    "regexp"
    "github.com/coderianx/gosugar"
)

// E-mail doğrulama
func EmailValidator() gosugar.Validator {
    pattern := regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)
    return func(s string) error {
        if !pattern.MatchString(s) {
            return fmt.Errorf("invalid email format")
        }
        return nil
    }
}

// Şifre gücü
func StrongPassword() gosugar.Validator {
    return func(s string) error {
        hasUpper := false
        hasLower := false
        hasDigit := false

        for _, ch := range s {
            if ch >= 'A' && ch <= 'Z' {
                hasUpper = true
            } else if ch >= 'a' && ch <= 'z' {
                hasLower = true
            } else if ch >= '0' && ch <= '9' {
                hasDigit = true
            }
        }

        if !hasUpper || !hasLower || !hasDigit {
            return fmt.Errorf("password must contain upper, lower, and digit")
        }
        return nil
    }
}

func main() {
    password := gosugar.Input(
        "Password: ",
        gosugar.NotEmpty(),
        gosugar.MinLen(8),
        StrongPassword(),
    )

    fmt.Println("Password accepted:", password)
}
```

### Pattern 3: Enum-like Selection

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func OneOf(options ...string) gosugar.Validator {
    return func(s string) error {
        for _, opt := range options {
            if s == opt {
                return nil
            }
        }
        return fmt.Errorf("must be one of: %v", options)
    }
}

func main() {
    priority := gosugar.Input(
        "Priority (LOW/MEDIUM/HIGH): ",
        OneOf("LOW", "MEDIUM", "HIGH"),
    )

    level := gosugar.Input(
        "Level (BEGINNER/INTERMEDIATE/ADVANCED): ",
        OneOf("BEGINNER", "INTERMEDIATE", "ADVANCED"),
    )

    fmt.Printf("Priority: %s, Level: %s\n", priority, level)
}
```

---

## Error Recovery Pattern

### Pattern 1: Graceful Fallback

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func loadConfigFile() string {
    // Try config.json
    content, ok := gosugar.Try(func() string {
        return gosugar.ReadFile("config.json")
    })
    if ok {
        fmt.Println("Loaded from config.json")
        return content
    }

    // Try config.local.json
    content, ok = gosugar.Try(func() string {
        return gosugar.ReadFile("config.local.json")
    })
    if ok {
        fmt.Println("Loaded from config.local.json")
        return content
    }

    // Fallback to default
    fmt.Println("Using default config")
    return `{"port": 8080}`
}

func main() {
    config := loadConfigFile()
    fmt.Println("Config:", config)
}
```

### Pattern 2: Startup Validation

```go
package main

import (
    "fmt"
    "os"
    "github.com/coderianx/gosugar"
)

func validateStartup() {
    fmt.Println("Validating configuration...")

    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Fatal error:", r)
            os.Exit(1)
        }
    }()

    // All required variables
    apiKey := gosugar.MustEnv("API_KEY")
    dbURL := gosugar.MustEnv("DATABASE_URL")
    port := gosugar.EnvInt("PORT", 0)

    if port == 0 {
        panic("PORT must be > 0")
    }

    fmt.Printf("✓ All validations passed\n")
    fmt.Printf("✓ API Key: %s...\n", apiKey[:10])
    fmt.Printf("✓ Database: %s\n", dbURL)
    fmt.Printf("✓ Port: %d\n", port)
}

func main() {
    validateStartup()
    fmt.Println("Application starting...")
}
```

### Pattern 3: Non-Critical Operations

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func loadOptionalConfig() string {
    // Non-critical. Hata varsa default kullan.
    content, ok := gosugar.Try(func() string {
        return gosugar.ReadFile("features.json")
    })

    // Or fallback
    result := gosugar.Or(content, ok, "{}")
    fmt.Println("Features config:", result)

    return result
}

func main() {
    loadOptionalConfig()
}
```

---

## Test Data Generation

### Pattern 1: Mock User Data

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

type User struct {
    ID       int
    Name     string
    Email    string
    IsActive bool
    Score    float64
}

func generateTestUsers(count int) []User {
    users := make([]User, count)

    for i := 0; i < count; i++ {
        users[i] = User{
            ID:       i + 1,
            Name:     "User" + gosugar.RandString(5),
            Email:    "user" + fmt.Sprint(i) + "@test.com",
            IsActive: gosugar.RandBool(),
            Score:    gosugar.RandFloat(0.0, 100.0),
        }
    }

    return users
}

func main() {
    users := generateTestUsers(5)
    for _, u := range users {
        fmt.Printf("%d. %s (%s) - Score: %.1f\n", u.ID, u.Name, u.Email, u.Score)
    }
}
```

### Pattern 2: Random Selection

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func main() {
    // Rastgele renk seç
    colors := []string{"red", "green", "blue", "yellow", "purple"}
    selectedColor := gosugar.Choice(colors)

    // Rastgele öncelik
    priorities := []string{"LOW", "MEDIUM", "HIGH"}
    priority := gosugar.Choice(priorities)

    fmt.Printf("Random: Color=%s, Priority=%s\n", selectedColor, priority)
}
```

---

## CLI Application Pattern

### Pattern: Interactive Menu

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func showMenu() string {
    fmt.Println("\n=== Main Menu ===")
    fmt.Println("1. Add User")
    fmt.Println("2. List Users")
    fmt.Println("3. Settings")
    fmt.Println("4. Exit")

    choice := gosugar.Input(
        "Choose (1-4): ",
        gosugar.NotEmpty(),
    )

    return choice
}

func addUser() {
    fmt.Println("\n--- Add User ---")

    name := gosugar.Input(
        "Name: ",
        gosugar.NotEmpty(),
        gosugar.MinLen(2),
    )

    age := gosugar.InputInt("Age: ", 0)

    fmt.Printf("\nUser added: %s (Age %d)\n", name, age)
}

func listUsers() {
    fmt.Println("\n--- Users ---")
    fmt.Println("1. Alice (25)")
    fmt.Println("2. Bob (30)")
    fmt.Println("3. Charlie (28)")
}

func main() {
    gosugar.EnvFile(".env")

    fmt.Println("Welcome to User Manager")

    for {
        choice := showMenu()

        switch choice {
        case "1":
            addUser()
        case "2":
            listUsers()
        case "3":
            fmt.Println("Settings TBD")
        case "4":
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Invalid choice")
        }
    }
}
```

---

## Best Practices

### 1. Başa doğrulama

```go
// ✅ Doğru
func init() {
    config := loadConfig()
    validateConfig(config)
    // Hata varsa panic, program durur
}

// ❌ Yanlış
func main() {
    // Hata gec catch edilir
}
```

### 2. Validatörler Reusable

```go
// ✅ Reusable
isEmail := func() gosugar.Validator {
    return func(s string) error {
        if !strings.Contains(s, "@") {
            return fmt.Errorf("invalid email")
        }
        return nil
    }
}

// Kullan
email1 := gosugar.Input("Email: ", isEmail())
email2 := gosugar.Input("Backup Email: ", isEmail())
```

### 3. Fallback Chain

```go
// ✅ Cascade fallbacks
loadFromEnv := func() (string, bool) {
    return gosugar.Try(func() string {
        return gosugar.MustEnv("CONFIG_PATH")
    })
}

loadFromFile := func() (string, bool) {
    return gosugar.Try(func() string {
        return gosugar.ReadFile("config.json")
    })
}

loadFromDefault := func() string {
    return `{"port": 8080}`
}

// Try'lardan birisi başarılı olursa kullan
v1, ok1 := loadFromEnv()
v2, ok2 := loadFromFile()

if ok1 {
    config = v1
} else if ok2 {
    config = v2
} else {
    config = loadFromDefault()
}
```

---

## İlgili Dosyalar

- [`../guides/getting-started.md`](../guides/getting-started.md) - Başlama rehberi
- [`../api/validators.md`](../api/validators.md) - Validatörler referansı

