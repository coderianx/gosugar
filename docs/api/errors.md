# API Reference: errors - Error Management

A module providing panic handling and error recovery mechanisms. Offers safe error handling patterns.

## 📋 Contents

- [Overview](#overview)
- [Functions](#functions)
- [Examples](#examples)
- [Patterns](#patterns)

---

## Overview

### Purpose

- Panic handling and error handling
- Try/catch-like safe execution
- Provide fallback values

### Key Features

- ✅ Generic type support
- ✅ Panic-safe execution (`Try`)
- ✅ Fallback mechanism (`Or`)
- ✅ Must pattern

---

## Functions

### 1. `Must[T any](v T, err error) T`

Panics if error occurs in functions returning `(T, error)`.

**Signature:**
```go
func Must[T any](v T, err error) T
```

**Type Parameter:**
- `T`: Return type

**Parameters:**
- `v` (T): Value to return on success
- `err` (error): Error (nil if successful)

**Return Value:**
- On success: `v`
- On error: panics

**Behavior:**
- If `err != nil` panics
- Otherwise returns `v`

**Example:**

```go
package main

import (
	"os"
	"github.com/coderianx/gosugar"
)

func main() {
	// Success (no error)
	file := gosugar.Must(os.Open("data.txt"))
	defer file.Close()

	// Panics if error
	// gosugar.Must(os.Open("nonexistent.txt")) // panic!
}
```

---

### 2. `Check(err error)`

For functions that only return error. Panics if error occurs.

**Signature:**
```go
func Check(err error)
```

**Parameters:**
- `err` (error): Error to check

**Behavior:**
- If `err != nil` panics

**Example:**

```go
package main

import (
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Success
	gosugar.Check(os.Mkdir("./data", 0755))

	// Panics if error
	// gosugar.Check(os.RemoveAll("/")) // panic!
}
```

---

### 3. `Try[T any](fn func() T) (T, bool)`

Recovers from panic. Safely runs code that might panic.

**Signature:**
```go
func Try[T any](fn func() T) (T, bool)
```

**Type Parameter:**
- `T`: Return type of the function

**Parameters:**
- `fn` (func() T): Function to execute

**Return Value:**
- `v` (T): Function's return value (zero-value if panic)
- `ok` (bool): `true` if successful, `false` if panic

**Behavior:**
- Executes `fn()`
- If panic occurs, recovers it and returns `ok=false`
- Otherwise returns `ok=true`

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Risky code
	value, ok := gosugar.Try(func() int {
		// This might panic
		return 100 / 0
	})

	if !ok {
		fmt.Println("Error: Code panicked")
	} else {
		fmt.Println("Success:", value)
	}
}
```

---

### 4. `Or[T any](v T, ok bool, fallback T) T`

Used with Try. Provides fallback value.

**Signature:**
```go
func Or[T any](v T, ok bool, fallback T) T
```

**Type Parameters:**
- `T`: Value type

**Parameters:**
- `v` (T): Main value
- `ok` (bool): Whether successful (`Try` return value)
- `fallback` (T): Value to use if unsuccessful

**Return Value:**
- If `ok=true`: `v`
- If `ok=false`: `fallback`

**Behavior:**
- Works like simple ternary operator

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	value, ok := gosugar.Try(func() int {
		return 100 / 0 // panic
	})

	result := gosugar.Or(value, ok, 0)
	fmt.Println("Result:", result) // 0 (fallback)
}
```

---

### 5. `Ignore(err error)`

Intentionally ignore an error.

**Signature:**
```go
func Ignore(err error)
```

**Parameters:**
- `err` (error): Error to ignore

**Behavior:**
- Ignores the error
- Useful to suppress linter warnings

**Example:**

```go
package main

import (
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Even if error, we don't care
	gosugar.Ignore(os.RemoveAll("./temp"))
}
```

---

## Examples

### Example 1: Must Pattern

```go
package main

import (
	"os"
	"github.com/coderianx/gosugar"
)

func main() {
	// Open file
	file := gosugar.Must(os.Open("config.json"))
	defer file.Close()

	// If successful, continue
	println("File opened")
}
```

### Example 2: Try/Or Pattern

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"strconv"
)

func main() {
	// Risky: convert string to integer
	value, ok := gosugar.Try(func() int {
		return gosugar.Must(strconv.Atoi("abc"))
	})

	// Use 0 if unsuccessful
	result := gosugar.Or(value, ok, 0)
	fmt.Println("Value:", result)
}
```

### Example 3: File Operations

```go
package main

import (
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Operations that might fail
	content, ok := gosugar.Try(func() string {
		return gosugar.ReadFile("data.txt")
	})

	if ok {
		println("Read:", content)
	} else {
		println("File couldn't be read, using default")
		content = "Default data"
	}
}
```

### Example 4: Custom Function with Try

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func divideByZero() int {
	return 100 / 0 // panic!
}

func main() {
	result, ok := gosugar.Try(divideByZero)
	if !ok {
		fmt.Println("Operation failed")
		result = 0
	}
	fmt.Println("Result:", result)
}
```

---

## Patterns

### Pattern 1: Startup Validation

Check required variables at startup:

```go
func main() {
	// Panic if error (can be caught at startup)
	port := gosugar.Must(strconv.Atoi(os.Getenv("PORT")))
	dbURL := gosugar.MustEnv("DATABASE_URL")
	
	// If successful, continue
	println("Port:", port)
}
```

### Pattern 2: Fallback Values

For optional operations:

```go
func main() {
	content, ok := gosugar.Try(func() string {
		return gosugar.ReadFile("config.json")
	})
	
	config := gosugar.Or(content, ok, "{}")
}
```

### Pattern 3: Multiple Try/Or

```go
func main() {
	// Try in sequence
	v1, ok1 := gosugar.Try(func() string {
		return gosugar.ReadFile("config.local.json")
	})
	
	v2, ok2 := gosugar.Try(func() string {
		return gosugar.ReadFile("config.json")
	})
	
	config := gosugar.Or(v1, ok1, gosugar.Or(v2, ok2, "{}"))
}
```

---

## Related Modules

- **`env.go`**: Environment variables (MustEnv)
- **`file.go`**: File operations (error handling)
- **`getting-started.md`**: Getting started guide
