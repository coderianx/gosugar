# API Reference: input - User Input

A module that provides functionality to get data from users via terminal. You can interactively get string, integer, float values.

## 📋 Contents

- [Overview](#overview)
- [Functions](#functions)
- [Examples](#examples)

---

## Overview

### Purpose

- Get data from user via terminal
- Validate inputs with validators
- Return default value on invalid input

### Key Features

- ✅ String, Integer, Float input
- ✅ Composable validators
- ✅ Default value support
- ✅ Automatic whitespace trimming
- ✅ Panic-based error handling

---

## Functions

### 1. `Input(prompt string, validators ...Validator) string`

Gets string input from user and validates with validators.

**Signature:**
```go
func Input(prompt string, validators ...Validator) string
```

**Parameters:**
- `prompt` (string): Question/guide text to display
- `validators` (variadic): Validators to apply (optional)

**Return Value:**
- User's input string (whitespace trimmed)

**Behavior:**
- Displays prompt and waits for input
- Cleans input with `strings.TrimSpace()`
- Runs each validator in sequence
- Panics if validation fails
- Returns value if validation succeeds

**Error Cases:**
- Validation error: `panic("invalid string input: ...")`
- Input read error: `panic("input error")`

**Example:**

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	// Simple input (no validator)
	name := gosugar.Input("Your name: ")
	println("Hello,", name)

	// Input with validators
	email := gosugar.Input(
		"Email: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(5),
	)
	println("Email:", email)
}
```

**Execution:**
```
Your name: John Doe
Hello, John Doe
Email: ab@test.com   # MinLen(5) error! Ask again
Email: valid@email.com
Email: valid@email.com
```

---

### 2. `InputInt(prompt string, defaultValue ...int) int`

Gets integer input from user. Returns default value if invalid input.

**Signature:**
```go
func InputInt(prompt string, defaultValue ...int) int
```

**Parameters:**
- `prompt` (string): Question to display
- `defaultValue` (variadic): Default value on invalid input

**Return Value:**
- Integer value (if successful) or default (if unsuccessful)

**Behavior:**
- Displays prompt
- Tries to convert to integer with `strconv.Atoi()`
- If successful: returns integer
- If unsuccessful:
  - If default value exists: returns it
  - If not: panics

**Error Cases:**
- Invalid format and no default: `panic("invalid integer input: ...")`

**Example:**

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	// Successful input
	age := gosugar.InputInt("Your age: ")
	println("Age:", age)

	// Invalid input, returns default
	port := gosugar.InputInt("Port (default 8080): ", 8080)
	println("Port:", port)
}
```

**Execution:**
```
Your age: abc       # Invalid, panic!
Port (default 8080): xyz
Port: 8080         # Returned default, no error
```

---

### 3. `InputFloat(prompt string, defaultValue ...float64) float64`

Gets float input from user. Returns default value if invalid input.

**Signature:**
```go
func InputFloat(prompt string, defaultValue ...float64) float64
```

**Parameters:**
- `prompt` (string): Question to display
- `defaultValue` (variadic): Default on invalid input

**Return Value:**
- Float64 value

**Behavior:**
- Displays prompt
- Tries to convert with `strconv.ParseFloat()`
- Success: returns float
- Failure: returns default or panics

**Example:**

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	price := gosugar.InputFloat("Price: ", 9.99)
	println("Price:", price)

	discount := gosugar.InputFloat("Discount rate (0-1): ")
	println("Discount:", discount)
}
```

---

## Examples

### Example 1: Simple Survey

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	fmt.Println("=== Survey Form ===\n")

	name := gosugar.Input(
		"Your name: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(2),
	)

	age := gosugar.InputInt("Your age: ", 0)

	email := gosugar.Input(
		"Email: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(5),
	)

	fmt.Printf("\nThank you %s! Your info has been saved.\n", name)
}
```

### Example 2: With Validators

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Username: 3-20 characters
	username := gosugar.Input(
		"Username (3-20 characters): ",
		gosugar.NotEmpty(),
		gosugar.MinLen(3),
		gosugar.MaxLen(20),
	)

	// Password: minimum 8 characters
	password := gosugar.Input(
		"Password (min 8 characters): ",
		gosugar.NotEmpty(),
		gosugar.MinLen(8),
	)

	fmt.Printf("Registration successful: %s\n", username)
}
```

### Example 3: Numeric Input

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	quantity := gosugar.InputInt(
		"Quantity: ",
		1,  // default: 1
	)

	price := gosugar.InputFloat(
		"Unit Price ($): ",
		0.0,  // default: 0
	)

	total := float64(quantity) * price
	fmt.Printf("Total: $%.2f\n", total)
}
```

---

## Related Modules

- **`validators.go`**: Validator types and built-in validators
- **`errors.go`**: Error handling
- **`env.go`**: Provide defaults from environment variables
