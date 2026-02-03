# GoSugar

A lightweight, **zero-dependency** Go utility library providing convenient helper functions for common tasks. Perfect for CLI applications, scripts, and small projects.

## ✨ Features

### Core Modules
- **Environment Management** (`env.go`) - Load `.env` files, read typed environment variables with defaults
- **User Input Handling** (`input.go`) - Read and validate terminal input (string, integer, float)
- **Input Validation** (`validators.go`) - Composable validators (NotEmpty, MinLen, MaxLen)
- **Random Data Generation** (`random.go`) - Generate random integers, floats, booleans, strings, and make choices
- **Error Handling** (`errors.go`) - Safe panic recovery, Must/Check patterns, fallback mechanisms
- **File Operations** (`file.go`) - Simple read, write, create, and append operations

## 📦 Installation

```bash
go get github.com/coderianx/gosugar
```

**Requirements:**
- Go 1.18+ (for generic support)
- No external dependencies (stdlib only)

## 🚀 Quick Start

### Load Environment Variables

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func main() {
    // Load from .env file
    gosugar.EnvFile(".env")
    
    // Read with defaults
    appName := gosugar.EnvString("APP_NAME", "MyApp")
    port := gosugar.EnvInt("PORT", 8080)
    debug := gosugar.EnvBool("DEBUG", false)
    
    // Required variable (panics if missing)
    apiKey := gosugar.MustEnv("API_KEY")
    
    fmt.Printf("%s running on port %d\n", appName, port)
}
```

### Get User Input

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func main() {
    // Simple input
    name := gosugar.Input("What is your name? ")
    
    // Input with validators
    email := gosugar.Input(
        "Email: ",
        gosugar.NotEmpty(),
        gosugar.MinLen(5),
    )
    
    // Typed input with fallback
    age := gosugar.InputInt("Age: ", 18)
    price := gosugar.InputFloat("Price: ", 9.99)
    
    fmt.Printf("User: %s (%d)\n", name, age)
}
```

### Random Data

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func main() {
    // Random integers
    dice := gosugar.RandInt(1, 6)
    
    // Random floats
    chance := gosugar.RandFloat(0.0, 1.0)
    
    // Random strings
    token := gosugar.RandString(32)
    
    // Random booleans
    flip := gosugar.RandBool()
    
    // Random selection from slice (type-safe generics)
    colors := []string{"red", "green", "blue"}
    chosen := gosugar.Choice(colors)
    
    fmt.Printf("Dice: %d, Token: %s, Color: %s\n", dice, token, chosen)
}
```

### File Operations

```go
package main

import "github.com/coderianx/gosugar"

func main() {
    // Read file
    content := gosugar.ReadFile("config.txt")
    
    // Write file (overwrites)
    gosugar.WriteFile("output.txt", "Hello World")
    
    // Create file (won't overwrite)
    gosugar.CreateFile("new.txt", "Initial content")
    
    // Append to file (creates if missing)
    gosugar.AppendFile("log.txt", "Log line\n")
}
```

### Error Handling

```go
package main

import (
    "os"
    "github.com/coderianx/gosugar"
)

func main() {
    // Must pattern: error → panic
    file := gosugar.Must(os.Open("config.json"))
    defer file.Close()
    
    // Check pattern: only check error
    gosugar.Check(os.Mkdir("data", 0755))
    
    // Try pattern: panic-safe execution
    result, ok := gosugar.Try(func() int {
        return 100 / 0 // This panics
    })
    
    // Or pattern: use fallback if error
    value := gosugar.Or(result, ok, 0) // Returns 0 if panic occurred
}
```

## 📚 API Documentation

### Environment (`env.go`)

```go
func EnvFile(path string)                              // Load .env file
func EnvString(key string, default ...string) string   // Read string with default
func EnvInt(key string, default ...int) int           // Read int with default
func EnvBool(key string, default ...bool) bool        // Read bool with default
func MustEnv(key string) string                        // Read required variable
```

### Input (`input.go`)

```go
func Input(prompt string, validators ...Validator) string        // Read string
func InputInt(prompt string, default ...int) int                 // Read int
func InputFloat(prompt string, default ...float64) float64       // Read float
```

### Validators (`validators.go`)

```go
func NotEmpty() Validator                  // Ensure non-empty
func MinLen(n int) Validator              // Minimum length
func MaxLen(n int) Validator              // Maximum length
```

### Random (`random.go`)

```go
func RandInt(min, max int) int            // Random int (inclusive)
func RandFloat(min, max float64) float64  // Random float
func RandBool() bool                      // Random boolean
func RandString(length int) string        // Random string
func Choice[T any](items []T) T          // Random element (generic)
```

### Errors (`errors.go`)

```go
func Must[T any](v T, err error) T       // Panic if error
func Check(err error)                     // Check error only
func Try[T any](fn func() T) (T, bool)   // Panic-safe execution
func Or[T any](v T, ok bool, fallback T) T  // Fallback on error
```

### Files (`file.go`)

```go
func ReadFile(path string) string              // Read file to string
func WriteFile(path, content string)           // Write file (overwrite)
func CreateFile(path, content string)          // Create file (safe)
func AppendFile(path, content string)          // Append to file
```

## 🏗️ Project Structure

```
gosugar/
├── env.go              # Environment variable management
├── input.go            # Terminal input handling
├── validators.go       # Input validation
├── random.go           # Random data generation
├── errors.go           # Error handling utilities
├── file.go             # File I/O operations
├── go.mod              # Module definition
├── test/
│   └── main.go         # Usage examples
├── README.md           # This file
├── info.md             # Detailed project analysis (Turkish)
└── ROADMAP.md          # Future features (Turkish)
```

## 📋 Examples

### Complete CLI Application

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func main() {
    // Load configuration
    gosugar.EnvFile(".env")
    appName := gosugar.EnvString("APP_NAME", "MyApp")
    
    fmt.Printf("Welcome to %s\n", appName)
    
    // Get user input
    username := gosugar.Input(
        "Username: ",
        gosugar.NotEmpty(),
        gosugar.MinLen(3),
        gosugar.MaxLen(20),
    )
    
    age := gosugar.InputInt("Age: ", 18)
    
    // Generate token
    token := gosugar.RandString(32)
    
    // Save to file
    profile := fmt.Sprintf("User: %s\nAge: %d\nToken: %s\n", username, age, token)
    gosugar.WriteFile("profile.txt", profile)
    
    fmt.Println("Profile saved!")
}
```

### Data Processing

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func main() {
    // Generate random test data
    for i := 0; i < 10; i++ {
        score := gosugar.RandInt(0, 100)
        passed := gosugar.RandBool()
        fmt.Printf("Test %d: Score=%d, Passed=%v\n", i+1, score, passed)
    }
    
    // Safe error handling
    content, ok := gosugar.Try(func() string {
        return gosugar.ReadFile("missing.txt")
    })
    
    result := gosugar.Or(content, ok, "No data available")
    fmt.Println(result)
}
```

## 🎯 Use Cases

✅ CLI applications and command-line tools
✅ Configuration management
✅ User input validation
✅ Testing and mock data generation
✅ File processing scripts
✅ Small utilities and helpers
✅ Learning Go (simple examples)

## ⚠️ Important Notes

### Panic Behavior
GoSugar uses `panic()` for error handling in most functions. This is suitable for:
- CLI applications
- Scripts and utilities
- Learning projects

For production systems, consider wrapping with error handling:
```go
defer func() {
    if r := recover(); r != nil {
        log.Fatalf("Error: %v", r)
    }
}()
```

### Zero Dependencies
GoSugar uses only the Go standard library. No external packages required.

## 🚀 Future Features

See [ROADMAP.md](ROADMAP.md) for planned features including:
- String manipulation utilities
- JSON serialization helpers
- Collection operations (Map, Filter, etc.)
- Logging system
- HTTP client wrapper
- Database helpers
- And 14 more features!

## 📄 Documentation

- **README.md** (this file) - Quick start and API overview
- **info.md** - Detailed project analysis (Turkish)
- **ROADMAP.md** - Planned features (Turkish)

## 💡 Design Principles

- **Simplicity First** - Wrap stdlib, don't replace it
- **Zero Dependencies** - Only Go stdlib used
- **Consistent API** - Familiar patterns for Go developers
- **Type Safety** - Use generics where beneficial
- **Educational** - Clean code for learning

## 📦 Module Info

- **Go Version**: 1.18+
- **Module Path**: `github.com/coderianx/gosugar`
- **License**: See repository

## 🤝 Contributing

Contributions welcome! Areas of interest:
- Bug fixes and improvements
- Additional validators
- More random data generators
- Extended file operations
- Performance optimizations

## 📝 License

See LICENSE file in repository for details.

---

**Happy coding with GoSugar! 🍬**
