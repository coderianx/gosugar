# API Reference: logger - Logging Utility

The `logger.go` module provides a simple, level-based logging system that helps you control application output with severity levels from Debug to Fatal.

## 📋 Contents

- [Overview](#overview)
- [Functions](#functions)
- [Log Levels](#log-levels)
- [Examples](#examples)
- [Design Decisions](#design-decisions)

---

## Overview

### Purpose

- Create structured log messages with timestamps
- Control log verbosity using log levels
- Support different severity levels for different message types
- Exit the application on fatal errors

### Key Features

- ✅ Five log levels (Debug, Info, Warn, Error, Fatal)
- ✅ Timestamp formatting for each log message
- ✅ Log level filtering (only messages at or above threshold are logged)
- ✅ Automatic program exit on Fatal level logs
- ✅ Simple and lightweight API

---

## Log Levels

Log levels control which messages are displayed. Lower levels are more verbose:

| Level | Constant | Value | Purpose |
|-------|----------|-------|---------|
| Debug | `DebugLevel` | 0 | Detailed diagnostic information for development |
| Info | `InfoLevel` | 1 | General informational messages |
| Warn | `WarnLevel` | 2 | Warning messages for potentially harmful situations |
| Error | `ErrorLevel` | 3 | Error messages for serious problems |
| Fatal | `FatalLevel` | 4 | Fatal errors that cause program exit |

**Example**: If you set the logger level to `WarnLevel`, only Warn, Error, and Fatal messages will be logged. Debug and Info messages will be ignored.

---

## Functions

### NewLogger

Creates a new Logger instance with the specified log level.

```go
func NewLogger(level LogLevel) *Logger
```

**Parameters:**
- `level` - The initial logging level threshold

**Returns:**
- A pointer to a new Logger instance

**Example:**
```go
logger := gosugar.NewLogger(gosugar.InfoLevel)
```

---

### SetLevel

Updates the logger's log level threshold at runtime.

```go
func (l *Logger) SetLevel(level LogLevel)
```

**Parameters:**
- `level` - The new logging level threshold

**Example:**
```go
logger.SetLevel(gosugar.DebugLevel) // Enable debug messages
```

---

### Debug

Logs a message at DEBUG level.

```go
func (l *Logger) Debug(msg string)
```

**Parameters:**
- `msg` - The message to log

**Example:**
```go
logger.Debug("Database connection established")
```

---

### Info

Logs a message at INFO level.

```go
func (l *Logger) Info(msg string)
```

**Parameters:**
- `msg` - The message to log

**Example:**
```go
logger.Info("Server started on port 8080")
```

---

### Warn

Logs a message at WARN level.

```go
func (l *Logger) Warn(msg string)
```

**Parameters:**
- `msg` - The message to log

**Example:**
```go
logger.Warn("Deprecated API endpoint used")
```

---

### Error

Logs a message at ERROR level.

```go
func (l *Logger) Error(msg string)
```

**Parameters:**
- `msg` - The message to log

**Example:**
```go
logger.Error("Failed to connect to database")
```

---

### Fatal

Logs a message at FATAL level and exits the program with code 1.

```go
func (l *Logger) Fatal(msg string)
```

**Parameters:**
- `msg` - The message to log

**Example:**
```go
logger.Fatal("Critical configuration error - shutting down")
// Program exits after this call
```

---

## Examples

### Basic Setup

```go
package main

import (
	"github.com/yourname/gosugar"
)

func main() {
	// Create logger at Info level
	logger := gosugar.NewLogger(gosugar.InfoLevel)

	logger.Debug("This won't be shown")  // Below threshold
	logger.Info("Application started")   // Shown: [2026-02-10 15:30:45] [INFO] Application started
	logger.Warn("Low memory")             // Shown: [2026-02-10 15:30:45] [WARN] Low memory
}
```

### Changing Log Level at Runtime

```go
logger := gosugar.NewLogger(gosugar.WarnLevel)

logger.Info("This is not shown")  // Below threshold

logger.SetLevel(gosugar.InfoLevel)
logger.Info("Now this is shown")  // At new threshold
```

### Using Different Log Levels

```go
logger := gosugar.NewLogger(gosugar.DebugLevel)

logger.Debug("Detailed diagnostic info")
logger.Info("General information")
logger.Warn("Warning about something")
logger.Error("An error occurred")
logger.Fatal("Critical error - exiting") // Program terminates
```

---

## Design Decisions

### Level-Based Filtering

Messages are only logged if their level is greater than or equal to the logger's current level. This allows you to:
- Run with `DebugLevel` during development to see everything
- Switch to `ErrorLevel` in production to see only critical issues

### Timestamp Formatting

Every log message includes a timestamp in `YYYY-MM-DD HH:MM:SS` format. This helps track:
- When events occurred
- Duration between events
- Time-based patterns in logs

### Fatal Exits Program

The `Fatal` method calls `os.Exit(1)` after logging. This ensures:
- Critical errors are logged before shutdown
- The program terminates with a non-zero exit code (indicating failure)
- Calling code doesn't need to handle fatal errors separately

### Simple String-Only API

Log methods accept only strings, not formatted arguments. This design choice:
- Keeps the API simple and lightweight
- Encourages using string building/formatting before logging
- Makes log output predictable and consistent

### Stdout Only

All messages go to standard output. For more complex logging needs:
- Redirect stdout to a file in your shell
- Use a wrapper around this logger
- Consider more advanced logging libraries
