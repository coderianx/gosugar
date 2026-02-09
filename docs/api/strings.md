# API Reference: strings - String Manipulation Utilities

A comprehensive module that provides helper functions for common string operations, case conversions, and text formatting.

## 📋 Contents

- [Overview](#overview)
- [Functions](#functions)
- [Examples](#examples)

---

## Overview

### Purpose

- Simplify common string operations
- Provide case conversion utilities
- Handle string padding and truncation
- Create URL-friendly slugs
- Reduce repetitive string manipulation code

### Key Features

- ✅ UTF-8 safe operations
- ✅ Case conversions (camelCase, snake_case, kebab-case, PascalCase)
- ✅ Padding and truncation
- ✅ String trimming and reversal
- ✅ Slug generation
- ✅ Easy-to-use wrappers around stdlib functions

---

## Functions

### 1. `Trim(s string) string`

Removes leading and trailing whitespace from a string.

**Signature:**
```go
func Trim(s string) string
```

**Parameters:**
- `s` (string): String to trim

**Return Value:**
- `string`: Trimmed string

**Behavior:**
- Removes all leading and trailing whitespace
- Keeps internal spaces unchanged
- Returns empty string if input is empty

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Trim("  hello world  ")
	fmt.Println(result) // "hello world"
	
	result = gosugar.Trim("\t\n  text  \n\t")
	fmt.Println(result) // "text"
}
```

---

### 2. `ToUpper(s string) string`

Converts a string to uppercase.

**Signature:**
```go
func ToUpper(s string) string
```

**Parameters:**
- `s` (string): String to convert

**Return Value:**
- `string`: Uppercase string

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.ToUpper("hello")
	fmt.Println(result) // "HELLO"
	
	result = gosugar.ToUpper("GoSugar")
	fmt.Println(result) // "GOSUGAR"
}
```

---

### 3. `ToLower(s string) string`

Converts a string to lowercase.

**Signature:**
```go
func ToLower(s string) string
```

**Parameters:**
- `s` (string): String to convert

**Return Value:**
- `string`: Lowercase string

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.ToLower("HELLO")
	fmt.Println(result) // "hello"
	
	result = gosugar.ToLower("GoSugar")
	fmt.Println(result) // "gosugar"
}
```

---

### 4. `Reverse(s string) string`

Reverses a string character by character (UTF-8 safe).

**Signature:**
```go
func Reverse(s string) string
```

**Parameters:**
- `s` (string): String to reverse

**Return Value:**
- `string`: Reversed string

**Behavior:**
- Correctly handles multi-byte UTF-8 characters
- Reverses emoji and special characters properly

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Reverse("hello")
	fmt.Println(result) // "olleh"
	
	result = gosugar.Reverse("😀🎉")
	fmt.Println(result) // "🎉😀"
}
```

---

### 5. `Contains(s, substr string) bool`

Checks if a string contains a substring.

**Signature:**
```go
func Contains(s, substr string) bool
```

**Parameters:**
- `s` (string): String to search in
- `substr` (string): Substring to search for

**Return Value:**
- `bool`: true if substring is found, false otherwise

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	found := gosugar.Contains("hello world", "world")
	fmt.Println(found) // true
	
	found = gosugar.Contains("hello", "xyz")
	fmt.Println(found) // false
}
```

---

### 6. `HasPrefix(s, prefix string) bool`

Checks if a string starts with a prefix.

**Signature:**
```go
func HasPrefix(s, prefix string) bool
```

**Parameters:**
- `s` (string): String to check
- `prefix` (string): Prefix to look for

**Return Value:**
- `bool`: true if starts with prefix, false otherwise

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	has := gosugar.HasPrefix("hello.go", "hello")
	fmt.Println(has) // true
	
	has = gosugar.HasPrefix("hello.go", ".go")
	fmt.Println(has) // false
}
```

---

### 7. `HasSuffix(s, suffix string) bool`

Checks if a string ends with a suffix.

**Signature:**
```go
func HasSuffix(s, suffix string) bool
```

**Parameters:**
- `s` (string): String to check
- `suffix` (string): Suffix to look for

**Return Value:**
- `bool`: true if ends with suffix, false otherwise

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	has := gosugar.HasSuffix("hello.go", ".go")
	fmt.Println(has) // true
	
	has = gosugar.HasSuffix("hello", ".go")
	fmt.Println(has) // false
}
```

---

### 8. `Replace(s, old, new string) string`

Replaces the first occurrence of old with new.

**Signature:**
```go
func Replace(s, old, new string) string
```

**Parameters:**
- `s` (string): String to process
- `old` (string): Substring to find
- `new` (string): Replacement substring

**Return Value:**
- `string`: String with first occurrence replaced

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Replace("hello hello", "hello", "hi")
	fmt.Println(result) // "hi hello"
	
	result = gosugar.Replace("cat and cat", "cat", "dog")
	fmt.Println(result) // "dog and cat"
}
```

---

### 9. `ReplaceAll(s, old, new string) string`

Replaces all occurrences of old with new.

**Signature:**
```go
func ReplaceAll(s, old, new string) string
```

**Parameters:**
- `s` (string): String to process
- `old` (string): Substring to find
- `new` (string): Replacement substring

**Return Value:**
- `string`: String with all occurrences replaced

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.ReplaceAll("hello hello", "hello", "hi")
	fmt.Println(result) // "hi hi"
	
	result = gosugar.ReplaceAll("cat and cat", "cat", "dog")
	fmt.Println(result) // "dog and dog"
}
```

---

### 10. `Split(s, sep string) []string`

Splits a string by a separator.

**Signature:**
```go
func Split(s, sep string) []string
```

**Parameters:**
- `s` (string): String to split
- `sep` (string): Separator string

**Return Value:**
- `[]string`: Slice of substrings

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	parts := gosugar.Split("a,b,c", ",")
	fmt.Println(parts) // ["a" "b" "c"]
	
	words := gosugar.Split("hello-world-go", "-")
	fmt.Println(words) // ["hello" "world" "go"]
}
```

---

### 11. `Join(strs []string, sep string) string`

Joins a slice of strings with a separator.

**Signature:**
```go
func Join(strs []string, sep string) string
```

**Parameters:**
- `strs` ([]string): Slice of strings to join
- `sep` (string): Separator string

**Return Value:**
- `string`: Joined string

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Join([]string{"a", "b", "c"}, ",")
	fmt.Println(result) // "a,b,c"
	
	result = gosugar.Join([]string{"hello", "world"}, " ")
	fmt.Println(result) // "hello world"
}
```

---

### 12. `Repeat(s string, count int) string`

Repeats a string multiple times.

**Signature:**
```go
func Repeat(s string, count int) string
```

**Parameters:**
- `s` (string): String to repeat
- `count` (int): Number of times to repeat

**Return Value:**
- `string`: Repeated string

**Panics:**
- If count is negative

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Repeat("ab", 3)
	fmt.Println(result) // "ababab"
	
	result = gosugar.Repeat("=", 10)
	fmt.Println(result) // "=========="
}
```

---

### 13. `PadLeft(s string, width int, char rune) string`

Pads a string on the left side to reach a specific width.

**Signature:**
```go
func PadLeft(s string, width int, char rune) string
```

**Parameters:**
- `s` (string): String to pad
- `width` (int): Target width
- `char` (rune): Character to pad with

**Return Value:**
- `string`: Left-padded string

**Behavior:**
- If string is already wider than width, returns unchanged
- Considers character count, not byte count

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.PadLeft("5", 3, '0')
	fmt.Println(result) // "005"
	
	result = gosugar.PadLeft("hello", 10, '-')
	fmt.Println(result) // "-----hello"
}
```

---

### 14. `PadRight(s string, width int, char rune) string`

Pads a string on the right side to reach a specific width.

**Signature:**
```go
func PadRight(s string, width int, char rune) string
```

**Parameters:**
- `s` (string): String to pad
- `width` (int): Target width
- `char` (rune): Character to pad with

**Return Value:**
- `string`: Right-padded string

**Behavior:**
- If string is already wider than width, returns unchanged
- Considers character count, not byte count

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.PadRight("5", 3, '0')
	fmt.Println(result) // "500"
	
	result = gosugar.PadRight("hello", 10, '-')
	fmt.Println(result) // "hello-----"
}
```

---

### 15. `Truncate(s string, length int, suffix string) string`

Truncates a string to a maximum length and adds a suffix if truncated.

**Signature:**
```go
func Truncate(s string, length int, suffix string) string
```

**Parameters:**
- `s` (string): String to truncate
- `length` (int): Maximum length (including suffix)
- `suffix` (string): Suffix to add if truncated (typically "...")

**Return Value:**
- `string`: Truncated string with suffix if needed

**Behavior:**
- If string is shorter than or equal to length, returns unchanged
- Suffix is included in the length calculation

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Truncate("hello world", 8, "...")
	fmt.Println(result) // "hello..."
	
	result = gosugar.Truncate("hi", 8, "...")
	fmt.Println(result) // "hi"
}
```

---

### 16. `Capitalize(s string) string`

Capitalizes the first character of a string.

**Signature:**
```go
func Capitalize(s string) string
```

**Parameters:**
- `s` (string): String to capitalize

**Return Value:**
- `string`: String with first character capitalized

**Behavior:**
- Only affects the first character
- Rest of the string remains unchanged
- Returns empty string if input is empty

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Capitalize("hello")
	fmt.Println(result) // "Hello"
	
	result = gosugar.Capitalize("HELLO")
	fmt.Println(result) // "HELLO"
}
```

---

### 17. `Decapitalize(s string) string`

Decapitalizes the first character of a string.

**Signature:**
```go
func Decapitalize(s string) string
```

**Parameters:**
- `s` (string): String to decapitalize

**Return Value:**
- `string`: String with first character lowercased

**Behavior:**
- Only affects the first character
- Rest of the string remains unchanged
- Returns empty string if input is empty

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Decapitalize("Hello")
	fmt.Println(result) // "hello"
	
	result = gosugar.Decapitalize("HELLO")
	fmt.Println(result) // "hELLO"
}
```

---

### 18. `CamelCase(s string) string`

Converts a string to camelCase.

**Signature:**
```go
func CamelCase(s string) string
```

**Parameters:**
- `s` (string): String to convert

**Return Value:**
- `string`: camelCase version

**Behavior:**
- Removes spaces, hyphens, and underscores
- First word is lowercase, subsequent words are capitalized
- Handles multiple separators

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.CamelCase("hello world")
	fmt.Println(result) // "helloWorld"
	
	result = gosugar.CamelCase("hello-world")
	fmt.Println(result) // "helloWorld"
	
	result = gosugar.CamelCase("hello_world")
	fmt.Println(result) // "helloWorld"
}
```

---

### 19. `SnakeCase(s string) string`

Converts a string to snake_case.

**Signature:**
```go
func SnakeCase(s string) string
```

**Parameters:**
- `s` (string): String to convert

**Return Value:**
- `string`: snake_case version

**Behavior:**
- Converts to lowercase
- Inserts underscores before uppercase letters
- Replaces spaces and hyphens with underscores
- Removes duplicate underscores

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.SnakeCase("hello world")
	fmt.Println(result) // "hello_world"
	
	result = gosugar.SnakeCase("HelloWorld")
	fmt.Println(result) // "hello_world"
	
	result = gosugar.SnakeCase("hello-world")
	fmt.Println(result) // "hello_world"
}
```

---

### 20. `KebabCase(s string) string`

Converts a string to kebab-case.

**Signature:**
```go
func KebabCase(s string) string
```

**Parameters:**
- `s` (string): String to convert

**Return Value:**
- `string`: kebab-case version

**Behavior:**
- Similar to snake_case but uses hyphens instead of underscores

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.KebabCase("hello world")
	fmt.Println(result) // "hello-world"
	
	result = gosugar.KebabCase("HelloWorld")
	fmt.Println(result) // "hello-world"
	
	result = gosugar.KebabCase("hello_world")
	fmt.Println(result) // "hello-world"
}
```

---

### 21. `PascalCase(s string) string`

Converts a string to PascalCase.

**Signature:**
```go
func PascalCase(s string) string
```

**Parameters:**
- `s` (string): String to convert

**Return Value:**
- `string`: PascalCase version

**Behavior:**
- Similar to camelCase but capitalizes the first letter
- Each word starts with an uppercase letter

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.PascalCase("hello world")
	fmt.Println(result) // "HelloWorld"
	
	result = gosugar.PascalCase("hello-world")
	fmt.Println(result) // "HelloWorld"
	
	result = gosugar.PascalCase("hello_world")
	fmt.Println(result) // "HelloWorld"
}
```

---

### 22. `Slugify(s string) string`

Converts a string to a URL-friendly slug.

**Signature:**
```go
func Slugify(s string) string
```

**Parameters:**
- `s` (string): String to convert

**Return Value:**
- `string`: URL-friendly slug

**Behavior:**
- Converts to lowercase
- Removes special characters
- Replaces spaces with hyphens
- Removes trailing hyphens
- Perfect for URLs and file names

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Slugify("Hello World!")
	fmt.Println(result) // "hello-world"
	
	result = gosugar.Slugify("GoSugar 1.0 Released!")
	fmt.Println(result) // "gosugar-10-released"
	
	result = gosugar.Slugify("  Contact Us  ")
	fmt.Println(result) // "contact-us"
}
```

---

## Examples

### Example 1: String Normalization

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// User input from form
	userInput := "  JoHn DoE  "
	
	// Normalize
	name := gosugar.Trim(userInput)
	name = gosugar.Capitalize(name)
	
	fmt.Printf("Welcome, %s!\n", name)
}
```

### Example 2: Generate Variable Names

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	titles := []string{
		"User Profile",
		"Product Details",
		"API Key",
	}
	
	for _, title := range titles {
		camel := gosugar.CamelCase(title)
		snake := gosugar.SnakeCase(title)
		pascal := gosugar.PascalCase(title)
		
		fmt.Printf("Title: %s\n", title)
		fmt.Printf("  camelCase: %s\n", camel)
		fmt.Printf("  snake_case: %s\n", snake)
		fmt.Printf("  PascalCase: %s\n", pascal)
		fmt.Println()
	}
}
```

### Example 3: URL Slug Generation

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	articleTitles := []string{
		"Getting Started with GoSugar",
		"Building CLI Applications",
		"Best Practices for Go",
	}
	
	for _, title := range articleTitles {
		slug := gosugar.Slugify(title)
		url := fmt.Sprintf("https://blog.example.com/articles/%s", slug)
		fmt.Println(url)
	}
}
```

### Example 4: Text Formatting

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Center text with padding
	text := "GoSugar"
	width := 20
	
	result := gosugar.PadLeft(text, (width+len(text))/2, ' ')
	result = gosugar.PadRight(result, width, ' ')
	
	fmt.Printf("|%s|\n", result)
	
	// Create separators
	separator := gosugar.Repeat("=", 40)
	fmt.Println(separator)
}
```

### Example 5: Table Column Formatting

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Table header
	headers := []string{"ID", "Name", "Status"}
	data := [][]string{
		{"1", "User1", "Active"},
		{"12", "LongUserName", "Inactive"},
		{"123", "Test", "Active"},
	}
	
	// Format columns
	for _, row := range data {
		for i, cell := range row {
			// Pad to column width
			padded := gosugar.PadRight(cell, 15, ' ')
			fmt.Print(padded)
		}
		fmt.Println()
	}
}
```

---

## Related Modules

- **`errors.go`**: Error handling
- **`input.go`**: User input with validation
- **`validators.go`**: Input validation helpers
