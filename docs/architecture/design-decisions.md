# Design Decisions - Why Was It Designed This Way?

This documentation explains **why** GoSugar was designed the way it was. Each decision includes trade-offs and alternatives.

## 📋 Contents

- [Why Panic?](#why-panic)
- [Why Single Package?](#why-single-package)
- [Why Generics?](#why-generics)
- [Why Function Types?](#why-function-types)
- [Why Zero Dependencies?](#why-zero-dependencies)
- [Frequently Asked Questions](#frequently-asked-questions)

---

## Why Panic?

### Decision: Error handling uses panic

**Code Example:**
```go
// GoSugar
func MustEnv(key string) string {
    value, ok := os.LookupEnv(key)
    if !ok || value == "" {
        panic(fmt.Errorf("required env var missing: %s", key))  // ← PANIC
    }
    return value
}
```

### Rationale?

| Situation | With Panic | With Error |
|-----------|-----------|-----------|
| **Config error** | ✅ Learn immediately | ❌ Keep running (wrong) |
| **Startup validation** | ✅ Clean | ❌ Check everywhere |
| **Code clarity** | ✅ Simple | ❌ Messy error handling |
| **Simple apps** | ✅ Appropriate | ❌ Unnecessary overhead |

### Decision Making

```
Target Audience: CLI apps, scripts, small projects
           ↓
Decision: Panic is appropriate
           ↓
Rationale: Config errors must be caught early
```

### Alternatives

**Alternative 1: Returning Errors**

```go
// ❌ More verbose
func EnvString(key string) (string, error) {
    value, ok := os.LookupEnv(key)
    if !ok {
        return "", fmt.Errorf("missing: %s", key)
    }
    return value, nil
}

// Usage
value, err := gosugar.EnvString("KEY")
if err != nil {
    // ... error handling
}
```

**Takeaway:** For production-grade error handling, `Try/Or` patterns available.

---

## Why Single Package?

### Decision: All functions in `gosugar` package

**Structure:**
```
gosugar/
├── env.go
├── input.go
├── validators.go
├── random.go
├── errors.go
├── file.go
└── http.go

# NOTE: no subdirectories
```

**Usage:**
```go
import "github.com/coderianx/gosugar"

gosugar.EnvString("KEY")
gosugar.Input("Q: ")
gosugar.RandInt(1, 10)
// All in one place
```

### Rationale?

| Aspect | Single Package | Multiple Packages |
|--------|---|---|
| **Import** | `import "...gosugar"` | Multiple imports needed |
| **Namespace** | `gosugar.Func()` | `env.Func()`, `input.Func()` |
| **Discovery** | ✅ All in one place | ❌ Where is it? |
| **Dependencies** | 🟡 Load all modules | ✅ Load only needed |
| **Simple apps** | ✅ Minimal imports | ❌ Multiple imports |

### Alternatives

**Alternative 1: Subpackages**

```
gosugar/
├── env/
│   └── env.go
├── input/
│   └── input.go
└── go.mod
```

**Usage:**
```go
import (
    "github.com/coderianx/gosugar/env"
    "github.com/coderianx/gosugar/input"
)

env.EnvString(...)
input.Input(...)
```

**Pros:**
- ✅ Optional imports
- ✅ Namespace organization

**Cons:**
- ❌ More complex
- ❌ Where do validators go? (shared?)

**Alternative 2: Monolithic File**

```
gosugar.go (1000+ lines)
```

**Cons:**
- ❌ Hard to read
- ❌ Track dependencies difficult

---

## Why Generics?

### Decision: Using Go 1.18+ generics

**Code Example:**
```go
// ✅ GoSugar - Generic
func Choice[T any](items []T) T {
    return items[rand.Intn(len(items))]
}

// Usage
fruit := gosugar.Choice([]string{"apple", "banana"})
num := gosugar.Choice([]int{1, 2, 3})
```

### Rationale?

| Advantage | Explanation |
|-----------|-------------|
| **Type Safety** | Compile-time checking. Less human error. |
| **No Casting** | No `interface{}` conversion needed |
| **Performance** | No runtime reflection. Fast. |
| **Clarity** | Intent is clear. Easy to read. |

**Example: Without Generics**

```go
// ❌ Go 1.17 - interface{}
func Choice(items []interface{}) interface{} {
    return items[rand.Intn(len(items))]
}

// Usage
data := gosugar.Choice([]interface{}{"a", "b"})
fruit := data.(string)  // ← Type assertion needed
```

### Alternatives

**Alternative 1: Type-Specific Functions**

```go
// ❌ Repetition
func ChoiceString(items []string) string { ... }
func ChoiceInt(items []int) int { ... }
func ChoiceFloat(items []float64) float64 { ... }
```

**Cons:**
- ❌ Too many functions
- ❌ Violates DRY principle

---

## Why Function Types?

### Decision: Validators as function types

**Code Example:**
```go
// ✅ GoSugar - Function Type
type Validator func(string) error

func MinLen(n int) Validator {
    return func(s string) error {
        if len(s) < n {
            return fmt.Errorf("min %d", n)
        }
        return nil
    }
}
```

### What Is This?

**Functional Programming Pattern:**

```
MinLen(5) call:
    ↓
Returns a function (closure with captured 'n')
    ↓
Used in Input("Q: ", MinLen(5))
    ↓
Validator function runs for each input
```

### Rationale?

| Reason | Explanation |
|--------|-------------|
| **Composability** | Validators can be chained |
| **Flexibility** | Users can write custom validators |
| **Simplicity** | Simpler than interface |
| **Higher-Order Funcs** | FP pattern in modern Go |

**Example: Chaining**

```go
// Input runs each validator
Input(
    "Q: ",
    NotEmpty(),       // validator 1
    MinLen(5),        // validator 2
    MaxLen(100),      // validator 3
)
```

### Alternatives

**Alternative 1: Interface**

```go
// ❌ Over-engineered
type Validator interface {
    Validate(string) error
}

type NotEmptyValidator struct{}
func (n NotEmptyValidator) Validate(s string) error { ... }

type MinLenValidator struct{ n int }
func (m MinLenValidator) Validate(s string) error { ... }

// Usage
Input("Q: ", NotEmptyValidator{}, MinLenValidator{5})
```

**Cons:**
- ❌ Too much code
- ❌ Boilerplate

**Alternative 2: Struct with Methods**

```go
// ❌ More state
type InputValidator struct {
    NotEmpty bool
    MinLen   int
    MaxLen   int
}
```

---

## Why Zero Dependencies?

### Decision: Only Go stdlib

**go.mod:**
```go
module github.com/coderianx/gosugar

go 1.25.5

// No require statements!
```

### Rationale?

| Reason | Effects |
|--------|---------|
| **Simplicity** | Easy to start. No complex setup. |
| **Stability** | External package updates don't break things |
| **Size** | Small binary |
| **Production** | Minimal deployment risk |
| **Learning** | Learn stdlib. Pick up best practices. |

### Trade-offs

| Scenario | Zero Deps | With Deps |
|----------|-----------|-----------|
| **String manipulation** | ✅ stdlib enough | ❌ `github.com/urfave/cli` |
| **HTTP requests** | ✅ net/http | ❌ `github.com/go-resty/resty` |
| **JSON** | ✅ encoding/json | ❌ `github.com/json-iterator/go` |

---

## Frequently Asked Questions

### Q: Can I use GoSugar in production?

**A:** Partially:
- ✅ **General utilities:** `Input`, `RandInt`, `File` OK
- ✅ **Config management:** `env` OK
- ⚠️ **Error handling:** Panic too aggressive, use `Try/Or`
- ❌ **High-frequency ops:** No HTTP, Database

**Recommendation:**
```go
// Production: config management only
gosugar.EnvFile(".env")
port := gosugar.EnvInt("PORT", 8080)

// Non-critical: CLI input
name := gosugar.Input("Name: ")

// Critical: use stdlib
db, err := sql.Open(...)
if err != nil { /* proper error handling */ }
```

### Q: Why not Try/Or before panic?

**A:** UX perspective:
- Panic: **Error message clear**
- Try/Or: **Error silently ignored** (hard to notice)

```go
// Panic: clear
port := gosugar.EnvInt("PORT")  // CRASH, clear message

// Try/Or: silent
value, ok := gosugar.Try(func() int {
    return gosugar.EnvInt("PORT")
})
port := gosugar.Or(value, ok, 8080)  // Silently uses 8080
```

### Q: Why is HTTP module limited?

**A:** Proper HTTP is complex:
- Custom headers
- POST/PUT/DELETE
- Timeout
- Retry logic
- Authentication

**Decision:** Support simple GET calls, avoid full REST clients.

**Takeaway:** Use `net/http` package directly (better).

### Q: I have a special use-case. Can I extend?

**A:** **Yes!** Write your own validator, pattern:

```go
// Custom validator
func NumericOnly() gosugar.Validator {
    return func(s string) error {
        for _, ch := range s {
            if ch < '0' || ch > '9' {
                return fmt.Errorf("only numbers")
            }
        }
        return nil
    }
}

// Use
code := gosugar.Input("Code: ", NumericOnly())
```

### Q: Why English documentation?

**A:** To serve non-Turkish speaking users while maintaining consistency with Turkish originals.

---

## Conclusion

GoSugar design philosophy:

```
┌─────────────────────────────────────┐
│  Simplicity > Flexibility            │
│  Clarity > Performance (95% cases)   │
│  Single Package > Modular Packages   │
│  Type Safety > Dynamic              │
│  Zero Deps > Feature Completeness   │
└─────────────────────────────────────┘
```

**Suitable for:**
- ✅ CLI applications
- ✅ Scripts and automation
- ✅ Prototyping
- ✅ Learning Go

**Not suitable for:**
- ❌ Enterprise systems
- ❌ High-performance apps
- ❌ Complex business logic

---

## Related Files

- [`ARCHITECTURE.md`](ARCHITECTURE.md) - Technical architecture
- [`../guides/design-patterns.md`](../guides/design-patterns.md) - Usage patterns
