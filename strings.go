package gosugar

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// BASIC STRING OPERATIONS

// Trim removes leading and trailing whitespace from a string.
// It's a wrapper around strings.TrimSpace for convenience.
//
// Parameters:
//   - s: String to trim
//
// Returns:
//   - Trimmed string without leading/trailing whitespace
//
// Example:
//
//	Trim("  hello world  ") // "hello world"
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// ToUpper converts a string to uppercase.
// It's a wrapper around strings.ToUpper.
//
// Parameters:
//   - s: String to convert
//
// Returns:
//   - Uppercase version of the string
//
// Example:
//
//	ToUpper("hello") // "HELLO"
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower converts a string to lowercase.
// It's a wrapper around strings.ToLower.
//
// Parameters:
//   - s: String to convert
//
// Returns:
//   - Lowercase version of the string
//
// Example:
//
//	ToLower("HELLO") // "hello"
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Reverse reverses a string character by character.
// Handles multi-byte UTF-8 characters correctly.
//
// Parameters:
//   - s: String to reverse
//
// Returns:
//   - Reversed string
//
// Example:
//
//	Reverse("hello") // "olleh"
//	Reverse("😀🎉") // "🎉😀" (handles emoji correctly)
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// SEARCHING & CHECKING

// Contains checks if a string contains a substring.
// It's a wrapper around strings.Contains.
//
// Parameters:
//   - s: String to search in
//   - substr: Substring to search for
//
// Returns:
//   - true if substring is found, false otherwise
//
// Example:
//
//	Contains("hello world", "world") // true
//	Contains("hello", "xyz") // false
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// HasPrefix checks if a string starts with a prefix.
// It's a wrapper around strings.HasPrefix.
//
// Parameters:
//   - s: String to check
//   - prefix: Prefix to look for
//
// Returns:
//   - true if string starts with prefix, false otherwise
//
// Example:
//
//	HasPrefix("hello.go", ".go") // false
//	HasPrefix("hello.go", "hello") // true
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix checks if a string ends with a suffix.
// It's a wrapper around strings.HasSuffix.
//
// Parameters:
//   - s: String to check
//   - suffix: Suffix to look for
//
// Returns:
//   - true if string ends with suffix, false otherwise
//
// Example:
//
//	HasSuffix("hello.go", ".go") // true
//	HasSuffix("hello", ".go") // false
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// REPLACEMENT & MANIPULATION

// Replace replaces the first occurrence of old with new.
// It's a wrapper around strings.Replace with a count of 1.
//
// Parameters:
//   - s: String to process
//   - old: Substring to find
//   - new: Replacement substring
//
// Returns:
//   - String with first occurrence replaced
//
// Example:
//
//	Replace("hello hello", "hello", "hi") // "hi hello"
func Replace(s, old, new string) string {
	return strings.Replace(s, old, new, 1)
}

// ReplaceAll replaces all occurrences of old with new.
// It's a wrapper around strings.ReplaceAll.
//
// Parameters:
//   - s: String to process
//   - old: Substring to find
//   - new: Replacement substring
//
// Returns:
//   - String with all occurrences replaced
//
// Example:
//
//	ReplaceAll("hello hello", "hello", "hi") // "hi hi"
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// SPLITTING & JOINING

// Split splits a string by a separator.
// It's a wrapper around strings.Split.
//
// Parameters:
//   - s: String to split
//   - sep: Separator string
//
// Returns:
//   - Slice of substrings
//
// Example:
//
//	Split("a,b,c", ",") // ["a", "b", "c"]
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Join joins a slice of strings with a separator.
// It's a wrapper around strings.Join.
//
// Parameters:
//   - strs: Slice of strings to join
//   - sep: Separator string
//
// Returns:
//   - Joined string
//
// Example:
//
//	Join([]string{"a", "b", "c"}, ",") // "a,b,c"
func Join(strs []string, sep string) string {
	return strings.Join(strs, sep)
}

// REPETITION & PADDING

// Repeat repeats a string multiple times.
// It's a wrapper around strings.Repeat.
//
// Parameters:
//   - s: String to repeat
//   - count: Number of times to repeat
//
// Returns:
//   - Repeated string
//
// Panics:
//   - If count is negative
//
// Example:
//
//	Repeat("ab", 3) // "ababab"
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// PadLeft pads a string on the left side with a character to reach a specific width.
//
// Parameters:
//   - s: String to pad
//   - width: Target width
//   - char: Character to pad with
//
// Returns:
//   - Left-padded string
//
// Example:
//
//	PadLeft("5", 3, '0') // "005"
//	PadLeft("hello", 10, '-') // "-----hello"
func PadLeft(s string, width int, char rune) string {
	currentLen := utf8.RuneCountInString(s)
	if currentLen >= width {
		return s
	}
	padding := width - currentLen
	return Repeat(string(char), padding) + s
}

// PadRight pads a string on the right side with a character to reach a specific width.
//
// Parameters:
//   - s: String to pad
//   - width: Target width
//   - char: Character to pad with
//
// Returns:
//   - Right-padded string
//
// Example:
//
//	PadRight("5", 3, '0') // "500"
//	PadRight("hello", 10, '-') // "hello-----"
func PadRight(s string, width int, char rune) string {
	currentLen := utf8.RuneCountInString(s)
	if currentLen >= width {
		return s
	}
	padding := width - currentLen
	return s + Repeat(string(char), padding)
}

// TRUNCATION & FORMATTING

// Truncate truncates a string to a maximum length and adds a suffix.
// If the string is shorter than or equal to the length, it returns unchanged.
//
// Parameters:
//   - s: String to truncate
//   - length: Maximum length (including suffix)
//   - suffix: Suffix to add if truncated (typically "...")
//
// Returns:
//   - Truncated string with suffix if needed
//
// Example:
//
//	Truncate("hello world", 8, "...") // "hello..."
//	Truncate("hi", 8, "...") // "hi"
func Truncate(s string, length int, suffix string) string {
	if utf8.RuneCountInString(s) <= length {
		return s
	}
	runes := []rune(s)
	maxLen := length - utf8.RuneCountInString(suffix)
	if maxLen < 0 {
		maxLen = 0
	}
	return string(runes[:maxLen]) + suffix
}

// Capitalize capitalizes the first character of a string.
// The rest of the string remains unchanged.
//
// Parameters:
//   - s: String to capitalize
//
// Returns:
//   - String with first character capitalized
//
// Example:
//
//	Capitalize("hello") // "Hello"
//	Capitalize("HELLO") // "HELLO"
//	Capitalize("") // ""
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Decapitalize decapitalizes the first character of a string.
// The rest of the string remains unchanged.
//
// Parameters:
//   - s: String to decapitalize
//
// Returns:
//   - String with first character lowercased
//
// Example:
//
//	Decapitalize("Hello") // "hello"
//	Decapitalize("HELLO") // "hELLO"
//	Decapitalize("") // ""
func Decapitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// CASE CONVERSION

// CamelCase converts a string to camelCase.
// Removes spaces, hyphens, and underscores, capitalizing words after the first.
//
// Parameters:
//   - s: String to convert
//
// Returns:
//   - CamelCase version of the string
//
// Example:
//
//	CamelCase("hello world") // "helloWorld"
//	CamelCase("hello-world") // "helloWorld"
//	CamelCase("hello_world") // "helloWorld"
func CamelCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	// Replace common separators with spaces
	s = strings.NewReplacer(
		"-", " ",
		"_", " ",
		".", " ",
	).Replace(s)

	// Split by spaces
	parts := strings.Fields(s)
	if len(parts) == 0 {
		return ""
	}

	// First part is lowercase, rest are capitalized
	result := strings.ToLower(parts[0])
	for _, part := range parts[1:] {
		result += Capitalize(part)
	}

	return result
}

// SnakeCase converts a string to snake_case.
// Converts to lowercase and replaces spaces, hyphens with underscores.
//
// Parameters:
//   - s: String to convert
//
// Returns:
//   - snake_case version of the string
//
// Example:
//
//	SnakeCase("hello world") // "hello_world"
//	SnakeCase("HelloWorld") // "hello_world"
//	SnakeCase("hello-world") // "hello_world"
func SnakeCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	// Insert underscore before uppercase letters
	var result strings.Builder
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsUpper(r) && i > 0 && runes[i-1] != '_' && runes[i-1] != ' ' {
			result.WriteRune('_')
		}
		if r == ' ' || r == '-' {
			result.WriteRune('_')
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	return strings.ReplaceAll(result.String(), "__", "_")
}

// KebabCase converts a string to kebab-case.
// Similar to snake_case but uses hyphens instead of underscores.
//
// Parameters:
//   - s: String to convert
//
// Returns:
//   - kebab-case version of the string
//
// Example:
//
//	KebabCase("hello world") // "hello-world"
//	KebabCase("HelloWorld") // "hello-world"
//	KebabCase("hello_world") // "hello-world"
func KebabCase(s string) string {
	return strings.ReplaceAll(SnakeCase(s), "_", "-")
}

// PascalCase converts a string to PascalCase.
// Similar to CamelCase but capitalizes the first letter too.
//
// Parameters:
//   - s: String to convert
//
// Returns:
//   - PascalCase version of the string
//
// Example:
//
//	PascalCase("hello world") // "HelloWorld"
//	PascalCase("hello-world") // "HelloWorld"
//	PascalCase("hello_world") // "HelloWorld"
func PascalCase(s string) string {
	camel := CamelCase(s)
	if camel == "" {
		return camel
	}
	return Capitalize(camel)
}

// SLUGIFY

// Slugify converts a string to a URL-friendly slug.
// Converts to lowercase, removes special characters, replaces spaces with hyphens.
//
// Parameters:
//   - s: String to convert
//
// Returns:
//   - URL-friendly slug
//
// Example:
//
//	Slugify("Hello World!") // "hello-world"
//	Slugify("GoSugar 1.0 Released") // "gosugar-10-released"
func Slugify(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	var result strings.Builder
	lastWasDash := false

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(r)
			lastWasDash = false
		} else if unicode.IsSpace(r) || r == '-' || r == '_' {
			if !lastWasDash && result.Len() > 0 {
				result.WriteRune('-')
				lastWasDash = true
			}
		}
		// Skip other characters
	}

	// Remove trailing dash
	str := result.String()
	return strings.TrimSuffix(str, "-")
}