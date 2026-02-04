# API Reference: validators - Input Validation

A module that provides ready-to-use and composable validators for checking user inputs.

## 📋 Contents

- [Overview](#overview)
- [Validator Type](#validator-type)
- [Built-in Validators](#built-in-validators)
- [Writing Custom Validators](#writing-custom-validators)
- [Examples](#examples)

---

## Overview

### Purpose

- Validate user inputs
- Apply validators in chain fashion
- Enable writing custom validators

### Key Features

- ✅ Composable validators
- ✅ Functional programming pattern
- ✅ Error messages
- ✅ Extensible design

---

## Validator Type

### `Validator`

```go
type Validator func(string) error
```

**Description:**
- Takes a string parameter
- Returns `nil` if validation succeeds
- Returns `error` if validation fails

**Usage:**

```go
// Built-in validator
notEmpty := gosugar.NotEmpty()
err := notEmpty("")           // returns error
err = notEmpty("hello")       // returns nil

// Pass validator to Input
email := gosugar.Input(
	"Email: ",
	gosugar.NotEmpty(),  // This is Validator function
)
```

---

## Built-in Validators

### 1. `NotEmpty() Validator`

Rejects empty strings.

**Signature:**
```go
func NotEmpty() Validator
```

**Behavior:**
- If input is empty string: returns error ("value cannot be empty")
- If input is not empty: returns nil

**Example:**

```go
username := gosugar.Input(
	"Username: ",
	gosugar.NotEmpty(),
)
// Empty input causes "invalid string input" error
```

---

### 2. `MinLen(n int) Validator`

Checks minimum character count.

**Signature:**
```go
func MinLen(n int) Validator
```

**Parameters:**
- `n` (int): Minimum character count

**Behavior:**
- If `len(string) < n`: returns error
- For example: `len("hi") < 3` → error

**Error Message:**
```
"minimum length is 3"
```

**Example:**

```go
password := gosugar.Input(
	"Password (min 8): ",
	gosugar.NotEmpty(),
	gosugar.MinLen(8),
)
// "1234" input: MinLen(8) error
```

---

### 3. `MaxLen(n int) Validator`

Checks maximum character count.

**Signature:**
```go
func MaxLen(n int) Validator
```

**Parameters:**
- `n` (int): Maximum character count

**Behavior:**
- If `len(string) > n`: returns error

**Error Message:**
```
"maximum length is 100"
```

**Example:**

```go
bio := gosugar.Input(
	"Biography (max 200): ",
	gosugar.MaxLen(200),
)
```

---

## Writing Custom Validators

Since validators are functions, you can write your own:

### Pattern 1: Simple Validator

```go
package main

import "github.com/coderianx/gosugar"

// Numeric characters only
func NumericOnly() gosugar.Validator {
	return func(s string) error {
		for _, ch := range s {
			if ch < '0' || ch > '9' {
				return fmt.Errorf("contains non-numeric characters")
			}
		}
		return nil
	}
}

func main() {
	phoneNumber := gosugar.Input(
		"Phone: ",
		NumericOnly(),
		gosugar.MinLen(10),
	)
	println(phoneNumber)
}
```

### Pattern 2: Regex Validator

```go
package main

import (
	"fmt"
	"regexp"
	"github.com/coderianx/gosugar"
)

// Email format checking
func EmailFormat() gosugar.Validator {
	pattern := regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)
	return func(s string) error {
		if !pattern.MatchString(s) {
			return fmt.Errorf("invalid email format")
		}
		return nil
	}
}

func main() {
	email := gosugar.Input(
		"Email: ",
		gosugar.NotEmpty(),
		EmailFormat(),
	)
	println(email)
}
```

### Pattern 3: Parameterized Validator

```go
// Choose from "A", "B", "C"
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
	level := gosugar.Input(
		"Level (LOW/MEDIUM/HIGH): ",
		OneOf("LOW", "MEDIUM", "HIGH"),
	)
	println(level)
}
```

---

## Examples

### Example 1: Combination

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	username := gosugar.Input(
		"Username: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(3),
		gosugar.MaxLen(20),
	)
	println("Username:", username)
}
```

### Example 2: Different Validators

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Title: 5-100 characters
	title := gosugar.Input(
		"Title: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(5),
		gosugar.MaxLen(100),
	)

	// Description: 20-1000 characters
	description := gosugar.Input(
		"Description: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(20),
		gosugar.MaxLen(1000),
	)

	fmt.Println("Saved")
}
```

### Example 3: Custom + Built-in Validators

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"strings"
)

// Letters only
func LettersOnly() gosugar.Validator {
	return func(s string) error {
		for _, ch := range s {
			if !('a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z') {
				return fmt.Errorf("only letters allowed")
			}
		}
		return nil
	}
}

func main() {
	firstName := gosugar.Input(
		"First name: ",
		gosugar.NotEmpty(),
		LettersOnly(),
		gosugar.MinLen(2),
		gosugar.MaxLen(50),
	)

	lastName := gosugar.Input(
		"Last name: ",
		gosugar.NotEmpty(),
		LettersOnly(),
		gosugar.MinLen(2),
		gosugar.MaxLen(50),
	)

	fmt.Printf("Welcome, %s %s!\n", firstName, lastName)
}
```

---

## Related Modules

- **`input.go`**: Get input with validators
- **`errors.go`**: Error handling
- **`design-patterns.md`**: Custom validator writing examples
