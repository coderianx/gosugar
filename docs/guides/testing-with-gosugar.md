# Writing Tests with GoSugar

Learn how to use GoSugar when writing tests and creating test data.

## 📋 Contents

- [Test Data Generation](#test-data-generation)
- [Random Data Generators](#random-data-generators)
- [Test Fixtures](#test-fixtures)
- [Integration Testing](#integration-testing)
- [Examples](#examples)

---

## Test Data Generation

### Random Numbers

```go
package main

import (
    "testing"
    "github.com/coderianx/gosugar"
)

func TestWithRandomData(t *testing.T) {
    // Random between 1-100
    score := gosugar.RandInt(1, 100)
    
    if score < 1 || score > 100 {
        t.Errorf("Invalid score: %d", score)
    }
}

func TestMultipleRandoms(t *testing.T) {
    // Generate multiple randoms
    for i := 0; i < 10; i++ {
        val := gosugar.RandInt(0, 1000)
        if val < 0 || val > 1000 {
            t.Errorf("Out of range: %d", val)
        }
    }
}
```

### Random Strings

```go
package main

import (
    "testing"
    "github.com/coderianx/gosugar"
)

func TestTokenGeneration(t *testing.T) {
    // 32 character token
    token := gosugar.RandString(32)

    if len(token) != 32 {
        t.Errorf("Invalid token length: %d", len(token))
    }

    // Should only contain letters
    for _, ch := range token {
        if (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') {
            t.Errorf("Invalid character in token: %c", ch)
        }
    }
}
```

### Random Boolean

```go
package main

import (
    "testing"
    "github.com/coderianx/gosugar"
)

func TestBoolDistribution(t *testing.T) {
    // Generate random boolean 1000 times
    // Statistically expect ~500 true, ~500 false
    
    trueCount := 0
    for i := 0; i < 1000; i++ {
        if gosugar.RandBool() {
            trueCount++
        }
    }

    // Distribution check (400-600 range is normal)
    if trueCount < 400 || trueCount > 600 {
        t.Logf("Warning: Uneven distribution: %d/1000 true", trueCount)
    }
}
```

---

## Random Data Generators

### Faker-like Function

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

type TestUser struct {
    ID       int
    Username string
    Email    string
    IsActive bool
}

func GenerateTestUser() TestUser {
    id := gosugar.RandInt(1, 10000)
    username := fmt.Sprintf("user_%s", gosugar.RandString(8))
    email := fmt.Sprintf("%s@test.com", gosugar.RandString(6))
    active := gosugar.RandBool()

    return TestUser{
        ID:       id,
        Username: username,
        Email:    email,
        IsActive: active,
    }
}

func TestUserGeneration(t *testing.T) {
    // Multiple test users
    users := make([]TestUser, 5)
    for i := 0; i < 5; i++ {
        users[i] = GenerateTestUser()
    }

    // Verify uniqueness (rough check)
    ids := make(map[int]bool)
    for _, u := range users {
        if ids[u.ID] {
            t.Logf("Warning: Duplicate ID: %d", u.ID)
        }
        ids[u.ID] = true
    }
}
```

### Collection Generators

```go
package main

import (
    "fmt"
    "github.com/coderianx/gosugar"
)

func GenerateTestData(count int) []map[string]interface{} {
    data := make([]map[string]interface{}, count)

    options := []string{"A", "B", "C", "D"}

    for i := 0; i < count; i++ {
        data[i] = map[string]interface{}{
            "id":       i + 1,
            "name":     fmt.Sprintf("Item_%d", i),
            "value":    gosugar.RandInt(10, 100),
            "status":   gosugar.Choice(options),
            "active":   gosugar.RandBool(),
            "score":    gosugar.RandFloat(0.0, 1.0),
        }
    }

    return data
}
```

---

## Test Fixtures

### Fixture Files

```go
package main

import (
    "testing"
    "github.com/coderianx/gosugar"
)

func setupTestFiles(t *testing.T) {
    // Create test files
    testData := `{"id": 1, "name": "Test"}`
    
    gosugar.CreateFile("test_data.json", testData)
    
    t.Cleanup(func() {
        // Cleanup (Go 1.14+)
        // os.Remove("test_data.json")
    })
}

func TestWithFixture(t *testing.T) {
    setupTestFiles(t)

    // Read file
    content := gosugar.ReadFile("test_data.json")

    if len(content) == 0 {
        t.Error("Fixture file is empty")
    }
}
```

### Config Fixtures

```go
package main

import (
    "os"
    "testing"
    "github.com/coderianx/gosugar"
)

func setupTestEnv(t *testing.T) {
    // Setup test environment
    os.Setenv("TEST_MODE", "true")
    os.Setenv("DB_URL", "sqlite:///:memory:")
    os.Setenv("PORT", "9999")

    t.Cleanup(func() {
        os.Unsetenv("TEST_MODE")
        os.Unsetenv("DB_URL")
        os.Unsetenv("PORT")
    })
}

func TestWithEnvFixture(t *testing.T) {
    setupTestEnv(t)

    // Test env vars
    mode := gosugar.EnvString("TEST_MODE", "")
    if mode != "true" {
        t.Errorf("Expected TEST_MODE=true, got %s", mode)
    }

    port := gosugar.EnvInt("PORT", 0)
    if port != 9999 {
        t.Errorf("Expected PORT=9999, got %d", port)
    }
}
```

---

## Integration Testing

### File I/O Testing

```go
package main

import (
    "os"
    "testing"
    "github.com/coderianx/gosugar"
)

func TestFileOperations(t *testing.T) {
    tmpFile := "test_output.txt"
    defer os.Remove(tmpFile)

    // Test 1: Write
    testContent := "Hello, World!"
    gosugar.WriteFile(tmpFile, testContent)

    // Test 2: Read
    readContent := gosugar.ReadFile(tmpFile)
    if readContent != testContent {
        t.Errorf("Content mismatch: got %q, want %q", readContent, testContent)
    }

    // Test 3: Append
    appendContent := "\nAppended line"
    gosugar.AppendFile(tmpFile, appendContent)

    finalContent := gosugar.ReadFile(tmpFile)
    if finalContent != testContent+appendContent {
        t.Error("Append failed")
    }
}
```

### Input/Validator Testing

```go
package main

import (
    "testing"
    "github.com/coderianx/gosugar"
)

func TestValidators(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        valid   bool
        validators []gosugar.Validator
    }{
        {
            name:  "Empty string",
            input: "",
            valid: false,
            validators: []gosugar.Validator{gosugar.NotEmpty()},
        },
        {
            name:  "Min length",
            input: "abc",
            valid: false,
            validators: []gosugar.Validator{gosugar.MinLen(5)},
        },
        {
            name:  "Max length",
            input: "this is a very long string",
            valid: false,
            validators: []gosugar.Validator{gosugar.MaxLen(10)},
        },
        {
            name:  "Valid",
            input: "valid",
            valid: true,
            validators: []gosugar.Validator{
                gosugar.NotEmpty(),
                gosugar.MinLen(3),
                gosugar.MaxLen(10),
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var err error
            for _, v := range tt.validators {
                err = v(tt.input)
                if err != nil && tt.valid {
                    t.Errorf("Unexpected error: %v", err)
                }
                if err == nil && !tt.valid {
                    t.Error("Expected error but got none")
                }
            }
        })
    }
}
```

### Error Handling Testing

```go
package main

import (
    "testing"
    "github.com/coderianx/gosugar"
)

func TestErrorHandling(t *testing.T) {
    // Test Try/Or
    value, ok := gosugar.Try(func() int {
        return 42
    })

    if !ok {
        t.Error("Try should succeed")
    }

    if value != 42 {
        t.Errorf("Expected 42, got %d", value)
    }

    // Test Try with panic
    panicValue, panicOk := gosugar.Try(func() int {
        panic("test panic")
    })

    if panicOk {
        t.Error("Try should catch panic")
    }

    if panicValue != 0 {
        t.Errorf("Expected zero-value, got %d", panicValue)
    }
}
```

---

## Examples

### Example 1: Database Test with Random Data

```go
package main

import (
    "fmt"
    "testing"
    "github.com/coderianx/gosugar"
)

func insertTestRecords(count int) {
    for i := 0; i < count; i++ {
        record := map[string]interface{}{
            "id":     gosugar.RandInt(1, 1000000),
            "name":   fmt.Sprintf("User_%s", gosugar.RandString(5)),
            "score":  gosugar.RandFloat(0.0, 100.0),
            "active": gosugar.RandBool(),
        }

        // INSERT record into DB
        println(fmt.Sprintf("Inserted: %v", record))
    }
}

func TestDataInsertion(t *testing.T) {
    insertTestRecords(10)
}
```

### Example 2: Config Validation Test

```go
package main

import (
    "os"
    "testing"
    "github.com/coderianx/gosugar"
)

func TestConfigValidation(t *testing.T) {
    // Setup
    os.Setenv("APP_NAME", "TestApp")
    os.Setenv("PORT", "8080")
    os.Setenv("DEBUG", "true")

    defer func() {
        os.Unsetenv("APP_NAME")
        os.Unsetenv("PORT")
        os.Unsetenv("DEBUG")
    }()

    // Test
    name := gosugar.EnvString("APP_NAME")
    port := gosugar.EnvInt("PORT", 0)
    debug := gosugar.EnvBool("DEBUG", false)

    if name != "TestApp" {
        t.Errorf("APP_NAME mismatch")
    }

    if port != 8080 {
        t.Errorf("PORT mismatch")
    }

    if !debug {
        t.Errorf("DEBUG should be true")
    }
}
```

### Example 3: Parametrized Tests

```go
package main

import (
    "fmt"
    "testing"
    "github.com/coderianx/gosugar"
)

func TestRandomInRange(t *testing.T) {
    tests := []struct {
        min int
        max int
    }{
        {1, 10},
        {0, 100},
        {-10, 10},
    }

    for _, tt := range tests {
        t.Run(
            "RandInt_"+fmt.Sprintf("%d_%d", tt.min, tt.max),
            func(t *testing.T) {
                for i := 0; i < 100; i++ {
                    val := gosugar.RandInt(tt.min, tt.max)
                    if val < tt.min || val > tt.max {
                        t.Errorf("Out of range: %d (expected %d-%d)", val, tt.min, tt.max)
                    }
                }
            },
        )
    }
}
```

---

## Best Practices

| Rule | Description |
|------|-------------|
| **Deterministic Setup** | Use predetermined test data instead of random |
| **Cleanup** | Clean up test files with defer |
| **Isolation** | Each test should be independent |
| **Meaningful Names** | Test names should say what is tested |
| **Edge Cases** | Test empty, null, boundary values |

---

## Related Files

- [`../api/random.md`](../api/random.md) - Random API reference
- [`../api/file.md`](../api/file.md) - File API reference
