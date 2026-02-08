# API Reference: http - HTTP Request Operations

A module that provides functions for sending HTTP GET requests and receiving responses.

## 📋 Contents

- [Overview](#overview)
- [Functions](#functions)
- [Examples](#examples)

---

## Overview

### Purpose

- Send HTTP GET requests
- Send HTTP POST requests
- Read response body
- Decode JSON
- Get response headers

### Key Features

- ✅ GET requests
- ✅ POST requests
- ✅ JSON deserialization
- ✅ Header reading
- ✅ Error handling

---

## Functions

### 1. `GetBody(url string) (string, error)`

Makes an HTTP GET request and returns the body as a string.

**Signature:**
```go
func GetBody(url string) (string, error)
```

**Parameters:**
- `url` (string): URL to request

**Return Value:**
- `body` (string): Response body
- `error`: Error if occurs, nil otherwise

**Behavior:**
- Makes HTTP GET request
- Returns error if status code is not 200 OK
- Converts body to string and returns

**Error Cases:**
- Network error: returns error
- Non-200 status: `fmt.Errorf("status code: %d")`
- Body read error: returns error

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body, err := gosugar.GetBody("https://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Body:", body[:100]+"...")
}
```

---

### 2. `MustGetBody(url string) string`

Like `GetBody` but panics if error occurs.

**Signature:**
```go
func MustGetBody(url string) string
```

**Parameters:**
- `url` (string): URL to request

**Return Value:**
- `body` (string): Response body

**Behavior:**
- Calls GetBody
- Panics if error occurs
- Otherwise returns body

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Returns if successful
	body := gosugar.MustGetBody("https://httpbin.org/get")
	fmt.Println("Success:", body[:50])

	// Panics if error
	// body := gosugar.MustGetBody("https://invalid-url")
}
```

---

### 3. `GetJSON[T any](url string) (T, error)`

Makes an HTTP GET request and decodes the JSON response.

**Signature:**
```go
func GetJSON[T any](url string) (T, error)
```

**Type Parameter:**
- `T`: Struct type to decode into

**Parameters:**
- `url` (string): URL to request

**Return Value:**
- `result` (T): Decoded data
- `error`: Error if occurs, nil otherwise

**Behavior:**
- Calls GetBody
- Decodes JSON using `json.Unmarshal()`
- Returns value of type T on success

**Error Cases:**
- GetBody error: returns error
- JSON decode error: returns error

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	type Post struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	post, err := gosugar.GetJSON[Post]("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Post: %d - %s\n", post.ID, post.Title)
}
```

---

### 4. `GetHeader(url string) (http.Header, error)`

Makes an HTTP GET request and returns response headers.

**Signature:**
```go
func GetHeader(url string) (http.Header, error)
```

**Parameters:**
- `url` (string): URL to request

**Return Value:**
- `headers` (http.Header): Response headers
- `error`: Error if occurs

**Behavior:**
- Makes HTTP GET request
- Returns error if status code is not 200
- Returns headers
- `http.Header` is case-insensitive map

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	headers, err := gosugar.GetHeader("https://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Content-Type:", headers.Get("Content-Type"))
	fmt.Println("Server:", headers.Get("Server"))
}
```

---

### 5. `MustGetHeader(url string) http.Header`

Like `GetHeader` but panics if error occurs.

**Signature:**
```go
func MustGetHeader(url string) http.Header
```

**Parameters:**
- `url` (string): URL to request

**Return Value:**
- `headers` (http.Header): Response headers

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	headers := gosugar.MustGetHeader("https://example.com")
	fmt.Println("Content-Type:", headers.Get("Content-Type"))
}
```

---

### 6. `PostBody(url string, body io.Reader, contentType string) (string, error)`

Makes an HTTP POST request and returns the response body as a string.

**Signature:**
```go
func PostBody(url string, body io.Reader, contentType string) (string, error)
```

**Parameters:**
- `url` (string): URL to request
- `body` (io.Reader): POST request body
- `contentType` (string): Content-Type header value (e.g., "application/json")

**Return Value:**
- `body` (string): Response body
- `error`: Error if occurs, nil otherwise

**Behavior:**
- Makes HTTP POST request
- Returns error if status code is not 200 OK
- Converts body to string and returns

**Error Cases:**
- Network error: returns error
- Non-200 status: `fmt.Errorf("status code: %d")`
- Body read error: returns error

**Example:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte("test data"))
	
	response, err := gosugar.PostBody(
		"https://httpbin.org/post",
		body,
		"application/x-www-form-urlencoded",
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", response[:100])
}
```

---

### 7. `MustPostBody(url string, body io.Reader, contentType string) string`

Like `PostBody` but panics if error occurs.

**Signature:**
```go
func MustPostBody(url string, body io.Reader, contentType string) string
```

**Parameters:**
- `url` (string): URL to request
- `body` (io.Reader): POST request body
- `contentType` (string): Content-Type header value

**Return Value:**
- `body` (string): Response body

**Behavior:**
- Calls PostBody
- Panics if error occurs
- Otherwise returns response

**Example:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte("data"))
	
	response := gosugar.MustPostBody(
		"https://httpbin.org/post",
		body,
		"text/plain",
	)
	fmt.Println("Response:", response[:50])
}
```

---

### 8. `PostJSON[T any](url string, payload any) (T, error)`

Makes an HTTP POST request with JSON payload and decodes the JSON response.

**Signature:**
```go
func PostJSON[T any](url string, payload any) (T, error)
```

**Type Parameter:**
- `T`: Struct type to decode response into

**Parameters:**
- `url` (string): URL to request
- `payload` (any): Data to encode as JSON

**Return Value:**
- `result` (T): Decoded response data
- `error`: Error if occurs, nil otherwise

**Behavior:**
- Encodes payload to JSON
- Calls PostBody
- Decodes JSON response using `json.Unmarshal()`
- Returns value of type T on success

**Error Cases:**
- JSON encode error: returns error
- PostBody error: returns error
- JSON decode error: returns error

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	type CreatePostRequest struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
		UserID int    `json:"userId"`
	}

	type CreatePostResponse struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
		UserID int    `json:"userId"`
	}

	payload := CreatePostRequest{
		Title:  "New Post",
		Body:   "This is a new post",
		UserID: 1,
	}

	response, err := gosugar.PostJSON[CreatePostResponse](
		"https://jsonplaceholder.typicode.com/posts",
		payload,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Created Post: %d - %s\n", response.ID, response.Title)
}
```

---

### 9. `PostHeader(url string, body io.Reader, contentType string) (http.Header, error)`

Makes an HTTP POST request and returns response headers.

**Signature:**
```go
func PostHeader(url string, body io.Reader, contentType string) (http.Header, error)
```

**Parameters:**
- `url` (string): URL to request
- `body` (io.Reader): POST request body
- `contentType` (string): Content-Type header value

**Return Value:**
- `headers` (http.Header): Response headers
- `error`: Error if occurs

**Behavior:**
- Makes HTTP POST request
- Returns error if status code is not 200
- Returns headers
- `http.Header` is case-insensitive map

**Example:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte("test"))
	
	headers, err := gosugar.PostHeader(
		"https://httpbin.org/post",
		body,
		"text/plain",
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Content-Type:", headers.Get("Content-Type"))
	fmt.Println("Server:", headers.Get("Server"))
}
```

---

### 10. `MustPostHeader(url string, body io.Reader, contentType string) http.Header`

Like `PostHeader` but panics if error occurs.

**Signature:**
```go
func MustPostHeader(url string, body io.Reader, contentType string) http.Header
```

**Parameters:**
- `url` (string): URL to request
- `body` (io.Reader): POST request body
- `contentType` (string): Content-Type header value

**Return Value:**
- `headers` (http.Header): Response headers

**Example:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte("data"))
	
	headers := gosugar.MustPostHeader(
		"https://httpbin.org/post",
		body,
		"text/plain",
	)
	fmt.Println("Content-Type:", headers.Get("Content-Type"))
}
```

---

### 11. `PutBody(url string, body io.Reader, contentType string) (string, error)`

Makes an HTTP PUT request and returns the response body as a string.

**Signature:**
```go
func PutBody(url string, body io.Reader, contentType string) (string, error)
```

**Parameters:**
- `url` (string): URL to request
- `body` (io.Reader): PUT request body
- `contentType` (string): Content-Type header value (e.g., "application/json")

**Return Value:**
- `body` (string): Response body
- `error`: Error if occurs, nil otherwise

**Behavior:**
- Makes HTTP PUT request
- Returns error if status code is not 200 OK
- Converts body to string and returns

**Error Cases:**
- Network error: returns error
- Non-200 status: `fmt.Errorf("status code: %d")`
- Body read error: returns error

**Example:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte("updated data"))
	
	response, err := gosugar.PutBody(
		"https://httpbin.org/put",
		body,
		"application/json",
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", response[:100])
}
```

---

### 12. `MustPutBody(url string, body io.Reader, contentType string) string`

Like `PutBody` but panics if error occurs.

**Signature:**
```go
func MustPutBody(url string, body io.Reader, contentType string) string
```

**Parameters:**
- `url` (string): URL to request
- `body` (io.Reader): PUT request body
- `contentType` (string): Content-Type header value

**Return Value:**
- `body` (string): Response body

**Behavior:**
- Calls PutBody
- Panics if error occurs
- Otherwise returns response

**Example:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte("data"))
	
	response := gosugar.MustPutBody(
		"https://httpbin.org/put",
		body,
		"text/plain",
	)
	fmt.Println("Response:", response[:50])
}
```

---

### 13. `PutJSON[T any](url string, payload any) (T, error)`

Makes an HTTP PUT request with JSON payload and decodes the JSON response.

**Signature:**
```go
func PutJSON[T any](url string, payload any) (T, error)
```

**Type Parameter:**
- `T`: Struct type to decode response into

**Parameters:**
- `url` (string): URL to request
- `payload` (any): Data to encode as JSON

**Return Value:**
- `result` (T): Decoded response data
- `error`: Error if occurs, nil otherwise

**Behavior:**
- Encodes payload to JSON
- Calls PutBody
- Decodes JSON response using `json.Unmarshal()`
- Returns value of type T on success

**Error Cases:**
- JSON encode error: returns error
- PutBody error: returns error
- JSON decode error: returns error

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	type UpdatePostRequest struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	type UpdatePostResponse struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	payload := UpdatePostRequest{
		Title: "Updated Post",
		Body:  "This post has been updated",
	}

	response, err := gosugar.PutJSON[UpdatePostResponse](
		"https://jsonplaceholder.typicode.com/posts/1",
		payload,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Updated Post: %d - %s\n", response.ID, response.Title)
}
```

---

### 14. `PutHeader(url string, body io.Reader, contentType string) (http.Header, error)`

Makes an HTTP PUT request and returns response headers.

**Signature:**
```go
func PutHeader(url string, body io.Reader, contentType string) (http.Header, error)
```

**Parameters:**
- `url` (string): URL to request
- `body` (io.Reader): PUT request body
- `contentType` (string): Content-Type header value

**Return Value:**
- `headers` (http.Header): Response headers
- `error`: Error if occurs

**Behavior:**
- Makes HTTP PUT request
- Returns error if status code is not 200
- Returns headers
- `http.Header` is case-insensitive map

**Example:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte("test"))
	
	headers, err := gosugar.PutHeader(
		"https://httpbin.org/put",
		body,
		"text/plain",
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Content-Type:", headers.Get("Content-Type"))
	fmt.Println("Server:", headers.Get("Server"))
}
```

---

### 15. `MustPutHeader(url string, body io.Reader, contentType string) http.Header`

Like `PutHeader` but panics if error occurs.

**Signature:**
```go
func MustPutHeader(url string, body io.Reader, contentType string) http.Header
```

**Parameters:**
- `url` (string): URL to request
- `body` (io.Reader): PUT request body
- `contentType` (string): Content-Type header value

**Return Value:**
- `headers` (http.Header): Response headers

**Example:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte("data"))
	
	headers := gosugar.MustPutHeader(
		"https://httpbin.org/put",
		body,
		"text/plain",
	)
	fmt.Println("Content-Type:", headers.Get("Content-Type"))
}
```

---

### 16. `DeleteBody(url string, body io.Reader, contentType string) (string, error)`

Makes an HTTP DELETE request and returns the response body as a string.

**Signature:**
```go
func DeleteBody(url string, body io.Reader, contentType string) (string, error)
```

**Parameters:**
- `url` (string): URL to request
- `body` (io.Reader): DELETE request body (optional)
- `contentType` (string): Content-Type header value

**Return Value:**
- `body` (string): Response body
- `error`: Error if occurs, nil otherwise

**Behavior:**
- Makes HTTP DELETE request
- Returns error if status code is not 200 OK
- Converts body to string and returns

**Error Cases:**
- Network error: returns error
- Non-200 status: `fmt.Errorf("status code: %d")`
- Body read error: returns error

**Example:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte(""))
	
	response, err := gosugar.DeleteBody(
		"https://httpbin.org/delete",
		body,
		"application/json",
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", response[:100])
}
```

---

### 17. `MustDeleteBody(url string, body io.Reader, contentType string) string`

Like `DeleteBody` but panics if error occurs.

**Signature:**
```go
func MustDeleteBody(url string, body io.Reader, contentType string) string
```

**Parameters:**
- `url` (string): URL to request
- `body` (io.Reader): DELETE request body (optional)
- `contentType` (string): Content-Type header value

**Return Value:**
- `body` (string): Response body

**Behavior:**
- Calls DeleteBody
- Panics if error occurs
- Otherwise returns response

**Example:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte(""))
	
	response := gosugar.MustDeleteBody(
		"https://httpbin.org/delete",
		body,
		"text/plain",
	)
	fmt.Println("Response:", response[:50])
}
```

---

### 18. `DeleteJSON[T any](url string, payload any) (T, error)`

Makes an HTTP DELETE request with JSON payload and decodes the JSON response.

**Signature:**
```go
func DeleteJSON[T any](url string, payload any) (T, error)
```

**Type Parameter:**
- `T`: Struct type to decode response into

**Parameters:**
- `url` (string): URL to request
- `payload` (any): Data to encode as JSON

**Return Value:**
- `result` (T): Decoded response data
- `error`: Error if occurs, nil otherwise

**Behavior:**
- Encodes payload to JSON
- Calls DeleteBody
- Decodes JSON response using `json.Unmarshal()`
- Returns value of type T on success

**Error Cases:**
- JSON encode error: returns error
- DeleteBody error: returns error
- JSON decode error: returns error

**Example:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	type DeleteResponse struct {
		ID      int    `json:"id"`
		Deleted bool   `json:"deleted"`
	}

	payload := map[string]interface{}{
		"reason": "no longer needed",
	}

	response, err := gosugar.DeleteJSON[DeleteResponse](
		"https://jsonplaceholder.typicode.com/posts/1",
		payload,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Deleted: %v\n", response.Deleted)
}
```

---

### 19. `DeleteHeader(url string, body io.Reader, contentType string) (http.Header, error)`

Makes an HTTP DELETE request and returns response headers.

**Signature:**
```go
func DeleteHeader(url string, body io.Reader, contentType string) (http.Header, error)
```

**Parameters:**
- `url` (string): URL to request
- `body` (io.Reader): DELETE request body (optional)
- `contentType` (string): Content-Type header value

**Return Value:**
- `headers` (http.Header): Response headers
- `error`: Error if occurs

**Behavior:**
- Makes HTTP DELETE request
- Returns error if status code is not 200
- Returns headers
- `http.Header` is case-insensitive map

**Example:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte(""))
	
	headers, err := gosugar.DeleteHeader(
		"https://httpbin.org/delete",
		body,
		"text/plain",
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Content-Type:", headers.Get("Content-Type"))
	fmt.Println("Server:", headers.Get("Server"))
}
```

---

### 20. `MustDeleteHeader(url string, body io.Reader, contentType string) http.Header`

Like `DeleteHeader` but panics if error occurs.

**Signature:**
```go
func MustDeleteHeader(url string, body io.Reader, contentType string) http.Header
```

**Parameters:**
- `url` (string): URL to request
- `body` (io.Reader): DELETE request body (optional)
- `contentType` (string): Content-Type header value

**Return Value:**
- `headers` (http.Header): Response headers

**Example:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte(""))
	
	headers := gosugar.MustDeleteHeader(
		"https://httpbin.org/delete",
		body,
		"text/plain",
	)
	fmt.Println("Content-Type:", headers.Get("Content-Type"))
}
```

---

## Examples

### Example 1: API Call

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	type User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	// Get user from JSONPlaceholder test API
	user, err := gosugar.GetJSON[User](
		"https://jsonplaceholder.typicode.com/users/1",
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("User: %s (%s)\n", user.Name, user.Email)
}
```

### Example 2: Plain Text Response

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Get plain text response
	body, err := gosugar.GetBody("https://httpbin.org/html")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:")
	fmt.Println(body[:200])
}
```

### Example 3: HTTP Headers

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	headers, err := gosugar.GetHeader("https://github.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("=== GitHub Headers ===")
	fmt.Println("Server:", headers.Get("Server"))
	fmt.Println("Content-Type:", headers.Get("Content-Type"))
	fmt.Println("Connection:", headers.Get("Connection"))
}
```

---

## Limitations

⚠️ **Current Version Limitations:**

1. **Only 200 OK**: Other success status codes (2xx, 3xx) treated as errors
2. **No custom headers**: Cannot add Authorization etc.
3. **No timeout**: May wait indefinitely on slow connections
4. **No redirect handling**: Does not follow redirects
5. **No request cancellation**: Cannot cancel requests once started

**Workaround:** Use `net/http` package directly.

---

## Related Modules

- **`errors.go`**: Error handling
- **`getting-started.md`**: Getting started guide
