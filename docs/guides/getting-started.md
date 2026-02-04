# GoSugar - Getting Started Guide

This guide is prepared for you to take your first steps with GoSugar. You will learn everything from installation to writing your first program.

## 🎯 Goals of This Guide

After finishing this page, you will:
- ✅ Be able to install GoSugar
- ✅ Be able to use basic functions
- ✅ Be able to write your first CLI application
- ✅ Know where to look next (next steps)

**Reading time:** ~15 minutes

---

## 1️⃣ Installation

### Prerequisites
- Go 1.18 or higher
- Terminal/command line (bash, zsh, cmd, PowerShell, etc.)

### Installation Step

```bash
go get github.com/coderianx/gosugar
```

This command downloads the GoSugar library and adds it to your Go module.

### Verification

To verify the installation is successful, create a simple test file:

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Test: Random number
	num := gosugar.RandInt(1, 10)
	fmt.Println("Random:", num)
}
```

Run it:
```bash
go run main.go
```

If you see a number in the output, installation is complete! ✅

---

## 2️⃣ Basic Concepts

GoSugar consists of 6 basic modules:

### 📋 1. Environment Variables (`env`)
Your application's configuration (port, database URL, etc.)

```go
gosugar.EnvString("APP_NAME", "MyApp")  // OK: "MyApp" or env value
gosugar.EnvInt("PORT", 8080)            // OK: 8080 or env value
gosugar.MustEnv("API_KEY")              // Required: panics if missing
```

### ⌨️ 2. User Input (`input`)
Reading data from users in the terminal

```go
name := gosugar.Input("Your name: ")           // String
age := gosugar.InputInt("Your age: ", 18)   // Integer (default: 18)
price := gosugar.InputFloat("Price: ", 0)  // Float (default: 0)
```

### ✔️ 3. Validators (`validators`)
Validating inputs

```go
email := gosugar.Input(
	"Email: ",
	gosugar.NotEmpty(),    // Cannot be empty
	gosugar.MinLen(5),     // At least 5 characters
	gosugar.MaxLen(100),   // At most 100 characters
)
```

### 🎲 4. Random Data (`random`)
Random values for testing and demos

```go
dice := gosugar.RandInt(1, 6)              // Between 1-6
random := gosugar.RandString(10)           // 10 characters
options := []string{"A", "B", "C"}
choice := gosugar.Choice(options)          // Pick from list
```

### 🛡️ 5. Error Handling (`errors`)
Safe error handling

```go
file := gosugar.Must(os.Open("config.json"))       // Panic if error
gosugar.Check(someFunction())                       // Check error only
value, ok := gosugar.Try(riskyFunction)           // Safe execution
result := gosugar.Or(value, ok, defaultValue)    // With fallback
```

### 📁 6. File Operations (`file`)
File reading/writing

```go
content := gosugar.ReadFile("data.txt")         // Read
gosugar.WriteFile("output.txt", "Hello")       // Write
gosugar.CreateFile("new.txt", "Start")    // Create (skip if exists)
gosugar.AppendFile("log.txt", "Log line\n")  // Append
```

---

## 3️⃣ Your First Application

Now let's write a small but useful application: **Simple Survey Application**

### Step 1: Create a File

Create a new file called `survey.go`:

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	fmt.Println("=== Welcome ===\n")

	// Get user's name
	name := gosugar.Input(
		"Your name: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(2),
	)

	// Get age
	age := gosugar.InputInt("Your age: ", 0)

	// Get email
	email := gosugar.Input(
		"Email: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(5),
	)

	// Show results
	fmt.Println("\n=== Information You Entered ===")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Email: %s\n", email)
	fmt.Println("\nThank you!")
}
```

### Step 2: Run It

```bash
go run survey.go
```

### Step 3: Experience

```
=== Welcome ===

Your name: John
Your age: 25
Email: john@example.com

=== Information You Entered ===
Name: John
Age: 25
Email: john@example.com

Thank you!
```

**Congratulations!** You've written your first GoSugar application! 🎉

---

## 4️⃣ More Complex Example: Environment File

The core power of GoSugar is environment management. Let's see that:

### Step 1: Create a `.env` File

Create a file called `.env`:

```env
# Application Settings
APP_NAME=MyCLIApp
DEBUG=true
PORT=3000

# Database (example)
DB_HOST=localhost
DB_PORT=5432
```

### Step 2: Write Code

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Load .env file
	gosugar.EnvFile(".env")

	// Read environment variables
	appName := gosugar.EnvString("APP_NAME", "DefaultApp")
	debug := gosugar.EnvBool("DEBUG", false)
	port := gosugar.EnvInt("PORT", 8080)

	fmt.Printf("Application: %s\n", appName)
	fmt.Printf("Debug: %v\n", debug)
	fmt.Printf("Port: %d\n", port)

	// Optional variables
	theme := gosugar.EnvString("THEME", "dark")
	fmt.Printf("Theme: %s (default)\n", theme)
}
```

### Step 3: Run It

```bash
go run main.go
```

Output:
```
Application: MyCLIApp
Debug: true
Port: 3000
Theme: dark (default)
```

**Important:** Loading environment variables from `.env` file is secure for production and simplifies configuration management.

---

## 5️⃣ Example with Random Data: Mini Game

Let's write a simple "Guess the Number Game":

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	fmt.Println("🎮 Guess the Number Game")
	fmt.Println("Guess a number between 1-100!\n")

	// Pick a random number (1-100)
	secretNumber := gosugar.RandInt(1, 100)
	attempts := 0
	maxAttempts := 7

	for attempts < maxAttempts {
		attempts++

		// Get user's guess
		guess := gosugar.InputInt(
			fmt.Sprintf("Attempt %d/%d - Your guess: ", attempts, maxAttempts),
			0,
		)

		if guess == secretNumber {
			fmt.Printf("\n🎉 You got it! The number was %d. Success in %d attempts!\n", secretNumber, attempts)
			return
		} else if guess < secretNumber {
			fmt.Println("📈 Try a higher number")
		} else {
			fmt.Println("📉 Try a lower number")
		}
	}

	fmt.Printf("\n😢 Game over! The number was %d.\n", secretNumber)
}
```

Run it and play! 🎮

---

## 6️⃣ Error Handling Example

GoSugar manages errors with panic. Be careful when using this:

### Safe File Reading

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Method 1: Safe reading (with Try/Or)
	content, ok := gosugar.Try(func() string {
		return gosugar.ReadFile("config.json")
	})

	if !ok {
		fmt.Println("File could not be read, using default")
		content = "Default configuration"
	}

	fmt.Println(content)

	// Method 2: Simple reading (panics if file doesn't exist)
	// data := gosugar.ReadFile("important.json")
	// This code will panic if file doesn't exist!
}
```

---

## 7️⃣ Next Steps

You have successfully learned the basic concepts! Next steps:

### 📖 **For Deeper Learning**
- For specific modules: Read references in [`../api/`](../api/) folder
- To understand architecture: [`../architecture/ARCHITECTURE.md`](../architecture/ARCHITECTURE.md)

### 🛠️ **For Practice**
1. Write your own CLI application
2. Configure with `.env` file
3. Control input with validators
4. Try error handling

### 🚀 **For Special Scenarios**
- Error handling best practices: [`error-handling.md`](error-handling.md)
- Design patterns: [`design-patterns.md`](design-patterns.md)
- Writing tests: [`testing-with-gosugar.md`](testing-with-gosugar.md)

---

## ❓ Frequently Asked Questions

### Q: Can I use GoSugar for web applications?
**A:** Yes! Environment management, file operations, error handling work in web applications too. However, input reading is designed for CLI so it can't be used directly in web.

### Q: Which Go version is required?
**A:** Go 1.18+. GoSugar uses the generics feature (introduced in 1.18).

### Q: Are there external dependencies?
**A:** No! It only uses Go's standard library.

### Q: Is panic dangerous?
**A:** Not an issue in simple applications. For critical systems, use `Try/Or` to catch errors.

### Q: Can I write custom validators?
**A:** Yes! Just write a function: see [`design-patterns.md`](design-patterns.md).

---

## 🎓 Information

This guide does not cover:
- ❌ Basics of Go language (for loops, variables, etc.)
- ❌ Complete API of the library (see [`../api/`](../api/))
- ❌ Advanced usage (see [`../architecture/`](../architecture/))

---

**Ready for more reading?** Choose a module reference:
- 📋 [`../api/env.md`](../api/env.md) - Environment variables in depth
- ⌨️ [`../api/input.md`](../api/input.md) - User input in depth
- 🎲 [`../api/random.md`](../api/random.md) - Random data in depth
- 📁 [`../api/file.md`](../api/file.md) - File operations in depth

Or another guide:
- 🛡️ [`error-handling.md`](error-handling.md) - Error handling strategies
- 🏗️ [`design-patterns.md`](design-patterns.md) - Design patterns

Have questions? Open an issue on `github.com/coderianx/gosugar`! 🤝
