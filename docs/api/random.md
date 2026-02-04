# API Reference: random - Random Data Generation

A module that generates random numbers, strings, and selections. Useful for testing and demos.

## 📋 Contents

- [Overview](#overview)
- [Functions](#functions)
- [Examples](#examples)

---

## Overview

### Purpose

- Generate random integers, floats, booleans
- Create random strings
- Select random elements from lists

### Key Features

- ✅ Seed automatically initialized
- ✅ Type-safe generics
- ✅ Different ranges (inclusive/exclusive)
- ✅ Error validation

---

## Functions

### 1. `RandInt(min, max int) int`

Returns a random integer within specified range.

**Signature:**
```go
func RandInt(min, max int) int
```

**Parameters:**
- `min` (int): Minimum value (inclusive)
- `max` (int): Maximum value (inclusive)

**Return Value:**
- Random integer: `min <= x <= max`

**Behavior:**
- Panics if `min > max`
- Generates different random number each call

**Error Cases:**
- `min > max`: `panic("min cannot be greater than max")`

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Roll die (1-6)
	dice := gosugar.RandInt(1, 6)
	fmt.Println("Die:", dice)

	// Range 1-100
	num := gosugar.RandInt(1, 100)
	fmt.Println("Random:", num)

	// Negative numbers
	val := gosugar.RandInt(-10, 10)
	fmt.Println("Value:", val)
}
```

---

### 2. `RandFloat(min, max float64) float64`

Returns a random float within specified range.

**Signature:**
```go
func RandFloat(min, max float64) float64
```

**Parameters:**
- `min` (float64): Minimum value (inclusive)
- `max` (float64): Maximum value (exclusive)

**Return Value:**
- Random float64: `min <= x < max`

**Behavior:**
- Panics if `min >= max`
- **Maximum is exclusive** (0.0-1.0 range doesn't include 1.0)

**Error Cases:**
- `min >= max`: `panic("min must be less than max")`

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Range 0.0-1.0 (probability)
	chance := gosugar.RandFloat(0.0, 1.0)
	fmt.Printf("Chance: %.4f\n", chance)

	// Range 10.5-20.5
	price := gosugar.RandFloat(10.5, 20.5)
	fmt.Printf("Price: $%.2f\n", price)
}
```

---

### 3. `RandBool() bool`

Returns a random boolean value.

**Signature:**
```go
func RandBool() bool
```

**Return Value:**
- `true` or `false` (50/50 chance)

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	if gosugar.RandBool() {
		fmt.Println("Heads")
	} else {
		fmt.Println("Tails")
	}
}
```

---

### 4. `RandString(length int) string`

Returns a random string of specified length (letters only).

**Signature:**
```go
func RandString(length int) string
```

**Parameters:**
- `length` (int): String length

**Return Value:**
- Random string (A-Z and a-z characters)

**Behavior:**
- Only English letters (26 + 26 = 52 characters)
- Panics if `length <= 0`

**Error Cases:**
- `length <= 0`: `panic("length must be positive")`

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Create token (10 characters)
	token := gosugar.RandString(10)
	fmt.Println("Token:", token)

	// Create ID (32 characters)
	id := gosugar.RandString(32)
	fmt.Println("ID:", id)
}
```

---

### 5. `Choice[T any](items []T) T`

Selects a random element from a list.

**Signature:**
```go
func Choice[T any](items []T) T
```

**Type Parameter:**
- `T`: Any type (generic)

**Parameters:**
- `items` ([]T): Slice to choose from

**Return Value:**
- Randomly selected element

**Behavior:**
- Panics if slice is empty
- Type-safe (compile-time checking)

**Error Cases:**
- Empty slice: `panic("cannot choose from empty slice")`

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Choose from strings
	colors := []string{"red", "green", "blue", "yellow"}
	color := gosugar.Choice(colors)
	fmt.Println("Color:", color)

	// Choose from integers
	numbers := []int{10, 20, 30, 40, 50}
	num := gosugar.Choice(numbers)
	fmt.Println("Number:", num)

	// Choose from structs
	type User struct {
		Name string
	}
	users := []User{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Charlie"},
	}
	selected := gosugar.Choice(users)
	fmt.Println("Selected:", selected.Name)
}
```

---

## Examples

### Example 1: Game

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	fmt.Println("🎮 Dice Game")
	fmt.Println("Rolling 3 dice...\n")

	total := 0
	for i := 1; i <= 3; i++ {
		dice := gosugar.RandInt(1, 6)
		fmt.Printf("Die %d: %d\n", i, dice)
		total += dice
	}

	fmt.Printf("\nTotal: %d\n", total)
}
```

### Example 2: Random Selection

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Choose random day
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	day := gosugar.Choice(days)
	fmt.Println("Random day:", day)

	// Choose random priority
	priorities := []string{"LOW", "MEDIUM", "HIGH"}
	priority := gosugar.Choice(priorities)
	fmt.Println("Priority:", priority)
}
```

### Example 3: Random Token/ID

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// API token
	token := gosugar.RandString(32)
	fmt.Println("Token:", token)

	// Password reset code
	code := gosugar.RandString(6)
	fmt.Println("Code:", code)

	// Session ID
	sessionID := gosugar.RandString(64)
	fmt.Println("Session:", sessionID)
}
```

### Example 4: Test Data

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	fmt.Println("=== Test Data Generation ===\n")

	for i := 1; i <= 5; i++ {
		name := gosugar.RandString(8)
		age := gosugar.RandInt(18, 65)
		active := gosugar.RandBool()
		score := gosugar.RandFloat(0.0, 100.0)

		fmt.Printf("User %d: %s, Age: %d, Active: %v, Score: %.1f\n",
			i, name, age, active, score)
	}
}
```

---

## Related Modules

- **`errors.go`**: Error handling
- **`getting-started.md`**: First steps
