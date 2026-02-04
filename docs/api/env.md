# API Reference: env - Environment Variables Management

Environment variables control your application's configuration (port, database URL, API keys, etc.). The `env.go` module makes it easy to read environment variables and load `.env` files.

## 📋 Contents

- [Overview](#overview)
- [Functions](#functions)
- [Examples](#examples)
- [Design Decisions](#design-decisions)

---

## Overview

### Purpose

- Load environment variables from `.env` files
- Read String, Integer, Boolean variables with specified type
- Provide default values
- Validate required variables

### Key Features

- ✅ `.env` file support (empty lines and comments ignored)
- ✅ Automatic type conversion (string → int, bool)
- ✅ Default value support
- ✅ Required variable validation
- ✅ Preserves existing environment variables (no override)
- ✅ Panic-based error handling (panic on failed reads)

### Workflow

```
Read .env file
    ↓
Parse key=value pairs
    ↓
Set environment variables
    ↓
Read with get functions (with defaults)
```

---

## Functions

### 1. `EnvFile(path string)`

Reads a `.env` file and loads all variables into the environment.

**Signature:**
```go
func EnvFile(path string)
```

**Parameters:**
- `path` (string): Path to the `.env` file

**Behavior:**
- Opens the file and reads line by line
- Ignores lines starting with `# ` (comments)
- Skips empty lines
- Parses `key=value` format
- **Does not override existing environment variables** (if already in environment, skips it)
- Panics if error occurs

**Error Cases:**
- File not found: `panic("cannot open env file: ...")`
- Invalid format (no `=`): `panic("invalid env line: ...")`
- Setenv failed: `panic("failed to set env ...")`

**Example `.env` file:**
```env
# Application Settings
APP_NAME=MyApp
PORT=8080
DEBUG=true

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=secret123
```

**Usage:**
```go
package main

import "github.com/coderianx/gosugar"

func main() {
	// Load .env file
	gosugar.EnvFile(".env")
	
	// Now you can read variables
	appName := gosugar.EnvString("APP_NAME")  // "MyApp"
	port := gosugar.EnvInt("PORT")            // 8080
}
```

**Related Functions:** `EnvString`, `EnvInt`, `EnvBool`, `MustEnv`

---

### 2. `EnvString(key string, defaultValue ...string) string`

Reads a string type environment variable.

**Signature:**
```go
func EnvString(key string, defaultValue ...string) string
```

**Parameters:**
- `key` (string): Name of the environment variable
- `defaultValue` (variadic): Default value to return if variable doesn't exist (optional)

**Return Value:**
- The variable's value (or: default value or empty string "")

**Behavior:**
- Checks if an environment variable with name `key` exists
- If exists and not empty, returns its value
- If not:
  - If default value provided, returns it
  - If not provided, returns empty string ""
- **Does not panic** (fault-tolerant)

**Error Cases:**
- No error cases (always returns successfully)

**Examples:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Scenario 1: Environment variable exists
	os.Setenv("APP_NAME", "MyApp")
	name := gosugar.EnvString("APP_NAME", "DefaultApp")
	fmt.Println(name) // "MyApp"

	// Scenario 2: Environment variable doesn't exist, use default
	theme := gosugar.EnvString("THEME", "dark")
	fmt.Println(theme) // "dark"

	// Scenario 3: Environment variable doesn't exist, no default
	lang := gosugar.EnvString("LANG")
	fmt.Println(lang) // "" (empty string)

	// Scenario 4: Environment variable is empty string
	os.Setenv("EMPTY", "")
	val := gosugar.EnvString("EMPTY", "default")
	fmt.Println(val) // "default" (empty string returns default)
}
```

**Related Functions:** `EnvInt`, `EnvBool`, `MustEnv`

---

### 3. `EnvInt(key string, defaultValue ...int) int`

Reads an integer type environment variable and performs type conversion.

**Signature:**
```go
func EnvInt(key string, defaultValue ...int) int
```

**Parameters:**
- `key` (string): Name of the environment variable
- `defaultValue` (variadic): Default value to return if conversion fails

**Return Value:**
- The converted integer value

**Behavior:**
- Reads the environment variable
- Converts to integer using `strconv.Atoi()`
- If conversion fails:
  - If default value exists, returns it
  - If not, panics
- If variable doesn't exist:
  - If default value exists, returns it
  - If not, panics

**Error Cases:**
- Invalid format (not an integer): `panic("invalid int env var ...")`
- Variable missing and no default: `panic("missing env var ...")`

**Examples:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Scenario 1: Valid integer
	os.Setenv("PORT", "8080")
	port := gosugar.EnvInt("PORT")
	fmt.Println(port) // 8080

	// Scenario 2: Doesn't exist, use default
	timeout := gosugar.EnvInt("TIMEOUT", 30)
	fmt.Println(timeout) // 30

	// Scenario 3: Invalid format, use default
	os.Setenv("BAD_NUMBER", "not_a_number")
	value := gosugar.EnvInt("BAD_NUMBER", 0)
	fmt.Println(value) // 0

	// Scenario 4: Invalid and no default → panic!
	// num := gosugar.EnvInt("MISSING_VAR") // panic!
}
```

**Related Functions:** `EnvString`, `EnvBool`

---

### 4. `EnvBool(key string, defaultValue ...bool) bool`

Reads a boolean type environment variable. Supports various string formats.

**Signature:**
```go
func EnvBool(key string, defaultValue ...bool) bool
```

**Parameters:**
- `key` (string): Name of the environment variable
- `defaultValue` (variadic): Default value to return if conversion fails

**Return Value:**
- Boolean value (true or false)

**Supported Values:**
- **True**: `"true"`, `"1"`, `"yes"`, `"y"`, `"on"` (case-insensitive)
- **False**: `"false"`, `"0"`, `"no"`, `"n"`, `"off"` (case-insensitive)

**Behavior:**
- Reads the environment variable
- Checks value case-insensitively
- If not a recognized format:
  - If default value exists, returns it
  - If not, panics

**Error Cases:**
- Invalid format: `panic("invalid bool env var ...")`
- Variable missing and no default: `panic("missing env var ...")`

**Examples:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Scenario 1: Standard true/false
	os.Setenv("DEBUG", "true")
	debug := gosugar.EnvBool("DEBUG")
	fmt.Println(debug) // true

	os.Setenv("PRODUCTION", "false")
	prod := gosugar.EnvBool("PRODUCTION")
	fmt.Println(prod) // false

	// Scenario 2: Alternative true values
	os.Setenv("ENABLE_CACHE", "1")
	cache := gosugar.EnvBool("ENABLE_CACHE")
	fmt.Println(cache) // true

	os.Setenv("AUTO_START", "yes")
	auto := gosugar.EnvBool("AUTO_START")
	fmt.Println(auto) // true

	// Scenario 3: Alternative false values
	os.Setenv("SKIP_VALIDATION", "0")
	skip := gosugar.EnvBool("SKIP_VALIDATION")
	fmt.Println(skip) // false

	// Scenario 4: Default value
	verbose := gosugar.EnvBool("VERBOSE", false)
	fmt.Println(verbose) // false

	// Scenario 5: Invalid format, use default
	os.Setenv("INVALID_BOOL", "maybe")
	value := gosugar.EnvBool("INVALID_BOOL", true)
	fmt.Println(value) // true
}
```

**Related Functions:** `EnvString`, `EnvInt`

---

### 5. `MustEnv(key string) string`

Reads a **required** environment variable. Panics if variable doesn't exist.

**Signature:**
```go
func MustEnv(key string) string
```

**Parameters:**
- `key` (string): Name of the environment variable (required)

**Return Value:**
- The environment variable's value (string)

**Behavior:**
- Checks if an environment variable with name `key` exists
- If exists and not empty, returns its value
- If not or empty: **panics**
- Does not support default values (strictly required)

**Error Cases:**
- Variable missing: `panic("required env var missing: ...")`
- Variable empty: `panic("required env var missing: ...")`

**Use Cases:**
- API keys
- Database connection string
- Critical configuration values

**Examples:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Scenario 1: Variable exists
	os.Setenv("DATABASE_URL", "postgres://localhost/mydb")
	dbURL := gosugar.MustEnv("DATABASE_URL")
	fmt.Println(dbURL) // "postgres://localhost/mydb"

	// Scenario 2: Variable missing → panic!
	// apiKey := gosugar.MustEnv("API_KEY") // panic!

	// Scenario 3: Variable empty → panic!
	os.Setenv("EMPTY_VAR", "")
	// val := gosugar.MustEnv("EMPTY_VAR") // panic!
}
```

**Best Practice:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Check required variables at startup
	apiKey := gosugar.MustEnv("API_KEY")
	dbURL := gosugar.MustEnv("DATABASE_URL")
	
	// Use defaults for optional variables
	port := gosugar.EnvInt("PORT", 8080)
	debug := gosugar.EnvBool("DEBUG", false)
	
	fmt.Printf("Successfully loaded: API_KEY, DATABASE_URL, port=%d\n", port)
}
```

**Related Functions:** `EnvString`, `EnvInt`, `EnvBool`

---

## Examples

### Example 1: Simple Configuration

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Load .env file
	gosugar.EnvFile(".env")

	// Read configuration
	appName := gosugar.EnvString("APP_NAME", "MyApp")
	port := gosugar.EnvInt("PORT", 8080)
	debug := gosugar.EnvBool("DEBUG", false)

	fmt.Printf("App: %s\n", appName)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Debug: %v\n", debug)
}
```

`.env`:
```env
APP_NAME=ProductionApp
PORT=3000
DEBUG=false
```

### Example 2: Required and Optional Variables

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	gosugar.EnvFile(".env")

	// Required variables (panics if missing)
	databaseURL := gosugar.MustEnv("DATABASE_URL")
	apiKey := gosugar.MustEnv("API_KEY")

	// Optional variables (with defaults)
	logLevel := gosugar.EnvString("LOG_LEVEL", "info")
	maxConnections := gosugar.EnvInt("MAX_CONNECTIONS", 10)
	enableCache := gosugar.EnvBool("ENABLE_CACHE", true)

	fmt.Printf("Database: %s\n", databaseURL[:20]+"...")
	fmt.Printf("Log Level: %s\n", logLevel)
	fmt.Printf("Max Conn: %d\n", maxConnections)
	fmt.Printf("Cache: %v\n", enableCache)
}
```

### Example 3: Behavior Based on Environment

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	gosugar.EnvFile(".env")

	env := gosugar.EnvString("ENVIRONMENT", "development")

	switch env {
	case "production":
		// Strictly required variables
		_ = gosugar.MustEnv("DATABASE_URL")
		_ = gosugar.MustEnv("API_KEY")
		fmt.Println("Production mode: All required variables checked")

	case "development":
		// Used more flexibly
		dbURL := gosugar.EnvString("DATABASE_URL", "localhost:5432")
		fmt.Printf("Dev mode: Database = %s\n", dbURL)

	default:
		fmt.Println("Unknown environment")
		os.Exit(1)
	}
}
```

---

## Design Decisions

### 1. Why Panic Was Chosen?

`MustEnv` and `EnvInt`/`EnvBool` (without defaults) panic. Why?

**Reason:**
- Configuration errors should be caught early
- Used at application startup (startup validation)
- Better than running with wrong configuration

**Alternative:** Make it safe with `Try`:
```go
value, ok := gosugar.Try(func() string {
	return gosugar.MustEnv("CRITICAL_VAR")
})
```

### 2. Why Not Override Existing Environment Variables?

If a variable already exists in the environment, the value loaded from `.env` is not used. Why?

**Reason:**
- Environment variables can be set before application startup
- In Docker/Kubernetes containers, ENV variables are set when container starts
- .env file is only used as a "fallback"

**Result:**
```bash
# Start from command line
PORT=9000 go run main.go

# Even if EnvFile is called, PORT won't use 8080, uses 9000
port := gosugar.EnvInt("PORT", 8080) // 9000
```

### 3. Why Is Type Conversion Automatic?

Why have `EnvInt` and `EnvBool` separate instead of just `EnvString`?

**Reason:**
- Type safety: compile-time checking
- Error handling: conversion failures are caught
- Convenience: write `PORT=8080` in `.env` and use it directly as integer

### 4. Why Multiple Values for Boolean?

Why support `"true"`, `"1"`, `"yes"`, `"y"`, `"on"` for boolean?

**Reason:**
- Different cultures and tools use different formats
- Docker and `docker-compose` prefer `"1"`/`"0"`
- Human readable: `"yes"`/`"no"` feels more natural

**Result:**
```env
DEBUG=true        # Go style
CACHE_ENABLED=1   # Docker style
VERBOSE=yes       # Human readable
```

---

## Frequently Asked Questions

### Q: Can `.env` files be used in production?
**A:** Normally no. In production, environment variables are set from system environment (Docker ENV, Kubernetes secrets, system environment). `.env` is only for local development.

### Q: Can I load multiple `.env` files?
**A:** Yes, `EnvFile()` can be called multiple times:
```go
gosugar.EnvFile(".env")
gosugar.EnvFile(".env.local")  // Second file is loaded
```
However, first-defined values are preserved (no override).

### Q: How are empty lines and comments handled?
**A:** Empty lines (``) and lines starting with `#` are automatically skipped:
```env
# This is a comment

PORT=8080     # This is also a comment

# DISABLED=true (this is skipped)
```

### Q: What if `EnvFile` doesn't find the file?
**A:** It panics. To make it fault-tolerant:
```go
_, ok := gosugar.Try(func() {
	gosugar.EnvFile(".env.local")
})
```

### Q: How can I see the injected environment variables?
**A:** Use `os.Environ()`:
```go
for _, env := range os.Environ() {
	fmt.Println(env)
}
```

---

## Related Modules

- **`input.go`**: User input (default values can be set from env variables)
- **`errors.go`**: Panic and error handling
- **`file.go`**: File reading (used in EnvFile)

---

## Resources

- `env.go` source code
- [`getting-started.md`](../guides/getting-started.md) - Getting started guide
- [`design-decisions.md`](../architecture/design-decisions.md) - Design decisions
