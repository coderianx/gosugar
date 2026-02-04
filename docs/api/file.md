# API Reference: file - File Operations

A module that simplifies file reading, writing, creating, and appending operations.

## 📋 Contents

- [Overview](#overview)
- [Functions](#functions)
- [Examples](#examples)

---

## Overview

### Purpose

- Simplify file reading
- Ease file writing operations
- File creation and append operations

### Key Features

- ✅ UTF-8 string support
- ✅ Automatic error handling (panic)
- ✅ CreateFile preserves existing files
- ✅ AppendFile creates file if doesn't exist

---

## Functions

### 1. `ReadFile(path string) string`

Reads a file and returns its contents as a string.

**Signature:**
```go
func ReadFile(path string) string
```

**Parameters:**
- `path` (string): File path

**Return Value:**
- File contents (string)

**Behavior:**
- Uses `os.ReadFile()`
- Loads entire content into memory
- Panics if error occurs

**Error Cases:**
- File not found: `panic("cannot read file ...")`
- Permission error: `panic("cannot read file ...")`

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Read file
	content := gosugar.ReadFile("data.txt")
	fmt.Println(content)

	// Read JSON
	config := gosugar.ReadFile("config.json")
	fmt.Println("Config:", config)
}
```

---

### 2. `WriteFile(path string, content string)`

Writes to a file. Overwrites if exists, creates if doesn't.

**Signature:**
```go
func WriteFile(path string, content string)
```

**Parameters:**
- `path` (string): File path
- `content` (string): Content to write

**Behavior:**
- Uses `os.WriteFile()`
- Creates with 0644 permissions
- Completely replaces content if exists (not append!)
- Panics if error occurs

**Error Cases:**
- Permission error: `panic("cannot write file ...")`
- Invalid path: `panic("cannot write file ...")`

**Example:**

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	// Write new file
	gosugar.WriteFile("output.txt", "Hello World!")

	// Overwrite if exists
	gosugar.WriteFile("output.txt", "Updated content")
}
```

**Warning:** Previous content is deleted if file exists!

---

### 3. `CreateFile(path string, content string)`

Creates a file **only if it doesn't exist**. Does nothing if it exists.

**Signature:**
```go
func CreateFile(path string, content string)
```

**Parameters:**
- `path` (string): File path
- `content` (string): Initial content

**Behavior:**
- If file doesn't exist: creates and writes content
- If file exists: does nothing (exits silently)
- Creates with 0644 permissions

**Error Cases:**
- CreateFile failed: `panic("cannot create file ...")`
- Other errors: `panic("cannot check file ...")`

**Example:**

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	// First time: creates file
	gosugar.CreateFile("config.json", "{\"port\": 8080}")

	// Second time: does nothing
	gosugar.CreateFile("config.json", "{\"port\": 3000}")

	// Result: config.json still has {"port": 8080}
}
```

**Best Practice:** Ideal for template files, default configurations.

---

### 4. `AppendFile(path string, content string)`

**Appends** to a file. Creates if doesn't exist.

**Signature:**
```go
func AppendFile(path string, content string)
```

**Parameters:**
- `path` (string): File path
- `content` (string): Content to append

**Behavior:**
- If file exists: appends to end
- If file doesn't exist: creates and writes content
- Creates with 0644 permissions
- Doesn't modify existing content

**Error Cases:**
- Append failed: `panic("cannot append to file ...")`

**Example:**

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	// Write logs
	gosugar.AppendFile("app.log", "Server started\n")
	gosugar.AppendFile("app.log", "Connection established\n")
	gosugar.AppendFile("app.log", "User logged in\n")

	// Result: app.log contains all lines
}
```

---

## Examples

### Example 1: Configuration File

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Create default config (does nothing if exists)
	defaultConfig := `{
	"app_name": "MyApp",
	"port": 8080,
	"debug": false
}`
	gosugar.CreateFile("config.json", defaultConfig)

	// Read config
	config := gosugar.ReadFile("config.json")
	fmt.Println("Configuration:")
	fmt.Println(config)
}
```

### Example 2: Logging System

```go
package main

import (
	"fmt"
	"time"
	"github.com/coderianx/gosugar"
)

func main() {
	logFile := "app.log"

	// Logging function
	writeLog := func(level, message string) {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		entry := fmt.Sprintf("[%s] %s: %s\n", timestamp, level, message)
		gosugar.AppendFile(logFile, entry)
	}

	// Example logs
	writeLog("INFO", "Application started")
	writeLog("DEBUG", "Opening database connection")
	writeLog("INFO", "Database connection successful")
	writeLog("ERROR", "API key not found")

	// Read logs
	logs := gosugar.ReadFile(logFile)
	fmt.Println("=== Log ===")
	fmt.Println(logs)
}
```

### Example 3: Content Processing

```go
package main

import (
	"fmt"
	"strings"
	"github.com/coderianx/gosugar"
)

func main() {
	// Read file
	content := gosugar.ReadFile("input.txt")

	// Process
	lines := strings.Split(content, "\n")
	fmt.Printf("Line count: %d\n", len(lines))

	// Write
	result := strings.Join(lines, " ")
	gosugar.WriteFile("output.txt", result)

	fmt.Println("Processing complete")
}
```

### Example 4: Data Export

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Start export
	reportFile := "report.csv"

	// Header
	gosugar.CreateFile(reportFile, "ID,Name,Score\n")

	// Add data
	for i := 1; i <= 5; i++ {
		entry := fmt.Sprintf("%d,User%d,%.2f\n", i, i, float64(i)*10.5)
		gosugar.AppendFile(reportFile, entry)
	}

	// Read and display report
	report := gosugar.ReadFile(reportFile)
	fmt.Println("=== Report ===")
	fmt.Println(report)
}
```

---

## Best Practices

### 1. CreateFile for Template Files

```go
// Create default file on first run
defaultEnv := `APP_NAME=MyApp
PORT=8080
DEBUG=false`

gosugar.CreateFile(".env", defaultEnv)
```

### 2. AppendFile for Logging

```go
// Whenever logging is needed
gosugar.AppendFile("debug.log", "Operation started\n")
// ... operation is performed ...
gosugar.AppendFile("debug.log", "Operation completed\n")
```

### 3. WriteFile for Overwrite

```go
// If you want to completely replace content
newContent := processData(oldContent)
gosugar.WriteFile("processed.txt", newContent)
```

### 4. Error Handling (with Try/Or)

```go
// File might not exist and fail
content, ok := gosugar.Try(func() string {
	return gosugar.ReadFile("optional.txt")
})

data := gosugar.Or(content, ok, "default content")
```

---

## Related Modules

- **`errors.go`**: Error handling (Try/Or)
- **`env.go`**: Load .env file (EnvFile)
- **`getting-started.md`**: Getting started guide
