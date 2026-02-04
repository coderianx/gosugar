# GoSugar Architecture - In-Depth Explanation

This documentation explains the complete architecture, design decisions, and internal structure of the GoSugar library.

## 📋 Contents

- [General Architecture](#general-architecture)
- [Module Design](#module-design)
- [Data Flow](#data-flow)
- [Design Principles](#design-principles)
- [Dependency Graph](#dependency-graph)
- [Common Patterns](#common-patterns)

---

## General Architecture

### Single Package Design

GoSugar uses a **single-package** architecture:

```
github.com/coderianx/gosugar/
├── env.go              # Module 1
├── input.go            # Module 2
├── validators.go       # Module 3
├── random.go           # Module 4
├── errors.go           # Module 5
├── file.go             # Module 6
├── http.go             # Module 7
└── go.mod
```

**Advantages:**
- ✅ Simple import: `import "github.com/coderianx/gosugar"`
- ✅ Flat namespace: `gosugar.Input()`, `gosugar.RandInt()` etc.
- ✅ Easy discoverability
- ✅ Minimal files

**Disadvantages:**
- ❌ Namespace pollution when package grows
- ❌ Can't use modules independently
- ❌ Internal implementation details exposed

### Package Structure

```
package gosugar

// All public functions
func EnvString(...) string
func Input(...) string
func RandInt(...) int
// ... etc
```

---

## Module Design

### 1. Module: `env.go` (Environment Variables)

**Responsibility:**
- Load `.env` files
- Read environment variables (typed)
- Provide default values

**Dependencies:**
- Go stdlib: `os`, `bufio`, `fmt`, `strconv`, `strings`

**Functions:**
```
EnvFile(path)              → Load from .env
EnvString(key, default)    → Read string
EnvInt(key, default)       → Read int (type conversion)
EnvBool(key, default)      → Read bool (type conversion)
MustEnv(key)               → Read required
```

**Workflow:**

```
User calls EnvFile(".env")
    ↓
Open file
    ↓
Scan lines
    ↓
Parse "key=value"
    ↓
Skip comments (#) and empty lines
    ↓
os.Setenv() (only if not exists)
    ↓
Done

User calls EnvString("PORT", 8080)
    ↓
os.LookupEnv("PORT")
    ↓
If exists and not empty: return value
If not exists/empty: return default
```

**Design Decisions:**
- **Why not override existing variables?** In container environments (Docker) ENV variables are set at startup. `.env` is only fallback.
- **Why type conversion in functions?** Delegates to strconv, handles errors.
- **Why panic if MustEnv missing?** Config errors should be caught early.

---

### 2. Module: `input.go` (User Input)

**Responsibility:**
- Get user input from terminal
- Type conversion (string → int, float)
- Apply validators

**Dependencies:**
- Go stdlib: `bufio`, `fmt`, `os`, `strconv`, `strings`
- Internal: `validators.go` (Validator type)

**Functions:**
```
Input(prompt, validators...)     → String input
InputInt(prompt, default)        → Int input
InputFloat(prompt, default)      → Float input
inputRaw(prompt) → internal      → Raw string read
```

**Workflow:**

```
User calls Input("Name: ", NotEmpty(), MinLen(3))
    ↓
inputRaw() → bufio.Scanner read
    ↓
strings.TrimSpace()
    ↓
Run each validator
    ↓
Validation failed: panic(error)
Validation succeeded: return string
```

**Design Decisions:**
- **Why panic on validation error?** If input validation fails, user should retry.
- **Why separate Input/InputInt/InputFloat?** Type safety. Compile-time checking.

---

### 3. Module: `validators.go` (Validation)

**Responsibility:**
- Define Validator type
- Provide built-in validators
- Enable composable pattern

**Dependencies:**
- Go stdlib: `errors`, `fmt`

**Types and Functions:**
```
type Validator func(string) error    // Type definition

NotEmpty() Validator                 // Built-in
MinLen(n) Validator                  // Built-in
MaxLen(n) Validator                  // Built-in
```

**Functional Programming Pattern:**

```go
// Validator is a function type
type Validator func(string) error

// Returned function closes over 'n'
func MinLen(n int) Validator {
    return func(s string) error {
        if len(s) < n {
            return fmt.Errorf("minimum length is %d", n)
        }
        return nil
    }
}

// Chaining usage
validators := []Validator{
    NotEmpty(),
    MinLen(5),
    MaxLen(100),
}
```

**Design Decisions:**
- **Why function type?** Composable and extensible. Users can write custom validators.
- **Why closure?** Parameters (n) are "embedded" in validator.

---

### 4. Module: `random.go` (Random Data)

**Responsibility:**
- Generate random numbers
- Create random strings
- Select from lists

**Dependencies:**
- Go stdlib: `math/rand`, `time`

**Functions:**
```
init() func                          → Seed initialization (auto)
RandInt(min, max) int               → Range [min, max]
RandFloat(min, max) float64         → Range [min, max)
RandBool() bool                     → 50/50 chance
RandString(length) string           → Letters only
Choice[T](items []T) T              → Select from list (generic)
```

**Design Decisions:**
- **Why init()?** Seed automatically initialized. Different random each run.
- **Why RandInt inclusive, RandFloat exclusive?** Follows Go stdlib pattern (math/rand.Intn exclusive, Float64 [0,1))
- **Why Choice generic?** Type-safe. Works with string, int, struct, etc.

---

### 5. Module: `errors.go` (Error Handling)

**Responsibility:**
- Panic patterns (Must, Check)
- Panic recovery (Try)
- Fallback mechanism (Or)

**Dependencies:**
- Go stdlib: (none directly, only built-in defer/recover)

**Functions:**
```
Must[T](v T, err) T                 → Panic if err
Check(err)                          → Panic if err
Try[T](fn func() T) (T, bool)       → Panic recovery
Or[T](v, ok, fallback) T            → Ternary-like
Ignore(err)                         → Ignore error
```

**Workflow:**

```
Must Pattern:
    file, err := os.Open("file.txt")
    f := gosugar.Must(file, err)    // Panic if err
    ↓
    if err != nil {
        panic(err)
    }
    return file

Try/Or Pattern:
    value, ok := gosugar.Try(func() int {
        return 100 / 0              // panic
    })
    ↓
    defer recover catches panic
    ↓
    ok = false, value = zero-value
    ↓
    result := gosugar.Or(value, ok, 0) // returns 0
```

**Design Decisions:**
- **Why panic?** Failed critical operations should stop the program.
- **Why Try/Or?** Non-critical operations need graceful fallback.
- **What is Ignore?** Suppress linter warnings: better than `_ = err`.

---

### 6. Module: `file.go` (File Operations)

**Responsibility:**
- Read files
- Write files
- Create files (protective)
- Append to files

**Dependencies:**
- Go stdlib: `fmt`, `os`

**Functions:**
```
ReadFile(path) string               → Read
WriteFile(path, content)            → Write (overwrite)
CreateFile(path, content)           → Create (skip if exists)
AppendFile(path, content)           → Append (create if missing)
```

**Design Decisions:**
- **Why CreateFile skips if exists?** To protect template files, default configs.
- **AppendFile creates if missing?** Very common in logging. No check on every call.

---

### 7. Module: `http.go` (HTTP Requests)

**Responsibility:**
- Make HTTP GET requests
- Read response body
- Decode JSON
- Read headers

**Dependencies:**
- Go stdlib: `encoding/json`, `fmt`, `io`, `net/http`

**Functions:**
```
GetBody(url) (string, error)               → Read body
MustGetBody(url) string                    → Read body (panic)
GetJSON[T](url) (T, error)                 → Decode JSON
GetHeader(url) (http.Header, error)        → Headers
MustGetHeader(url) http.Header             → Headers (panic)
```

---

## Data Flow

### Scenario 1: CLI Application

```
┌──────────────────────────────┐
│    Application Starts        │
└──────────────┬───────────────┘
               │
               ▼
    ┌─────────────────────┐
    │ env.go              │
    │ EnvFile(".env")     │
    │ EnvString(...)      │ ──→ environment vars
    └─────────────────────┘
               │
               ▼
    ┌─────────────────────┐
    │ input.go            │
    │ Input("Q: ")        │ ──→ validators (validators.go)
    │ InputInt(...)       │
    └─────────────────────┘
               │
               ▼
    ┌──────────────────────────┐
    │ Processing Logic         │
    │ (User code)              │
    │ - random data gen        │ ──→ random.go
    │ - file I/O               │ ──→ file.go
    │ - error handling         │ ──→ errors.go
    └──────────────────────────┘
               │
               ▼
    ┌─────────────────────┐
    │ file.go             │
    │ WriteFile(...)      │ ──→ Output file
    │ AppendFile(...)     │
    └─────────────────────┘
```

### Scenario 2: API Communication

```
┌──────────────────────────┐
│ API Code                 │
└──────────┬───────────────┘
           │
           ▼
    ┌──────────────────────┐
    │ http.go              │
    │ GetJSON[T](url)      │
    └──────────┬───────────┘
               │
               ▼
    ┌──────────────────────┐
    │ errors.go (implicit) │
    │ If HTTP error        │
    │ return error         │
    └──────────┬───────────┘
               │
               ▼
    ┌──────────────────────┐
    │ User Code            │
    │ Use Try/Or pattern   │
    │ provide fallback     │ ──→ errors.go
    └──────────────────────┘
```

---

## Design Principles

### 1. **Simplicity First**

Wrapping stdlib, not replacing:

```go
// ✅ GoSugar - wrapper
func EnvString(key, default) string {
    return os.LookupEnv(key) // wrap stdlib
}

// ❌ Replacement (too complex)
// custom environment variable system
```

### 2. **Zero Dependencies**

Only Go stdlib:

```go
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    // ... only stdlib
)

// ❌ NO external packages
// import "github.com/some/package"
```

### 3. **Panic-Based Error Handling**

For simple applications:

```go
// ✅ Simple apps
apiKey := gosugar.MustEnv("API_KEY") // panic if missing

// ❌ Production apps (caution needed)
value, ok := gosugar.Try(someRiskyOp) // recover from panic
```

### 4. **Type Safety (Generics)**

Using Go 1.18+ generics:

```go
// ✅ Type-safe
choice := gosugar.Choice([]string{"A", "B"}) // string
num := gosugar.Choice([]int{1, 2})            // int

// ❌ Type-unsafe (any conversion)
choice := someChoice([]interface{}{...})
```

### 5. **Composability**

Validators can be chained:

```go
// ✅ Chained validators
Input(
    "Email: ",
    NotEmpty(),
    MinLen(5),
    MaxLen(100),
)

// ❌ Manual validation everywhere
if email == "" { ... }
if len(email) < 5 { ... }
```

---

## Dependency Graph

```
input.go ──────┐
               ├──→ validators.go
               │
               ├──→ errors.go (implicit panics)
               │
               └──→ Go stdlib

env.go ────────┐
               ├──→ errors.go (panic)
               │
               └──→ Go stdlib

file.go ───────┐
               ├──→ errors.go (panic)
               │
               └──→ Go stdlib

random.go ─────→ Go stdlib (independent)

http.go ───────┐
               ├──→ errors.go (implicit via error handling)
               │
               └──→ Go stdlib

errors.go ─────→ Go stdlib (independent)

validators.go ─→ Go stdlib (independent)
```

**Depth:** Maximum 2 levels (no circular dependencies)

---

## Common Patterns

### Pattern 1: Startup Configuration

```go
func main() {
    // Load config
    gosugar.EnvFile(".env")
    
    // Validate critical vars
    apiKey := gosugar.MustEnv("API_KEY")      // panic if missing
    port := gosugar.EnvInt("PORT", 8080)      // with default
    
    // Start app
    startServer(port, apiKey)
}
```

### Pattern 2: User Input Loop

```go
for {
    // Get input with validators
    command := gosugar.Input(
        "> ",
        gosugar.NotEmpty(),
    )
    
    // Process
    handleCommand(command)
    
    // If validation fails, asks again
}
```

### Pattern 3: Graceful Fallback

```go
// Try main source
config, err := getConfigFromAPI()
if err == nil {
    // Success
    useConfig(config)
} else {
    // Fallback: local file
    config = gosugar.ReadFile("config.local.json")
    useConfig(config)
}

// Or using Try/Or
config, ok := gosugar.Try(getConfigFromAPI)
config = gosugar.Or(config, ok, defaultConfig)
```

### Pattern 4: Test Data Generation

```go
func generateTestData(count int) {
    for i := 0; i < count; i++ {
        user := User{
            ID:   gosugar.RandInt(1, 10000),
            Name: gosugar.RandString(10),
            Active: gosugar.RandBool(),
        }
        saveUser(user)
    }
}
```

---

## Extensibility

GoSugar's design encourages extension:

### Writing Custom Validator

```go
func IsEmail() gosugar.Validator {
    return func(s string) error {
        if !strings.Contains(s, "@") {
            return fmt.Errorf("not an email")
        }
        return nil
    }
}

// Use
email := gosugar.Input("Email: ", IsEmail())
```

### Custom Error Pattern

```go
// Instead of Try/Or
if err := riskyOperation(); err != nil {
    log.Printf("Operation failed: %v", err)
    handleError(err)
}
```

### Adding New Module

Add new module as `packagename/modulename.go`:

```go
// Example: string.go
package gosugar

func Reverse(s string) string {
    // implementation
}
```

---

## Conclusion

GoSugar architecture:
- 📦 **Single-package** design (simple)
- 🎯 **Focused** functions (each module does one thing)
- 🔗 **Minimal coupling** (modules are independent)
- 🛡️ **Panic-based** error handling (for simple apps)
- 💪 **Extensible** (custom validators, patterns, etc.)

Next: [`design-decisions.md`](design-decisions.md) for details on design decisions.
