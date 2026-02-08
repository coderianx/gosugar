# API Referansı: http - HTTP İstek İşlemleri

HTTP GET istekleri göndermek ve yanıt almak için fonksiyonlar sağlayan modül.

## 📋 İçindekiler

- [Genel Bakış](#genel-bakış)
- [Fonksiyonlar](#fonksiyonlar)
- [Örnekler](#örnekler)

---

## Genel Bakış

### Amaç

- HTTP GET istekleri göndermek
- HTTP POST istekleri göndermek
- Response body'sini okumak
- JSON decode etmek
- Response headers almak

### Başlıca Özellikler

- ✅ GET istekleri
- ✅ POST istekleri
- ✅ JSON deserialization
- ✅ Header okuma
- ✅ Error handling

---

## Fonksiyonlar

### 1. `GetBody(url string) (string, error)`

HTTP GET isteği yapar ve body'sini string olarak döner.

**Signature:**
```go
func GetBody(url string) (string, error)
```

**Parametreler:**
- `url` (string): İstek yapılacak URL

**Dönüş Değeri:**
- `body` (string): Response body
- `error`: Hata varsa error, yoksa nil

**Behavior:**
- HTTP GET request yapar
- Status code 200 OK değilse: error döner
- Body'sini string'e çevirip döner

**Hata Durumları:**
- Network hatası: error döner
- Non-200 status: `fmt.Errorf("status code: %d")`
- Body okuma hatası: error döner

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body, err := gosugar.GetBody("https://example.com")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Body:", body[:100]+"...")
}
```

---

### 2. `MustGetBody(url string) string`

`GetBody` gibi ama hata varsa panic atar.

**Signature:**
```go
func MustGetBody(url string) string
```

**Parametreler:**
- `url` (string): İstek yapılacak URL

**Dönüş Değeri:**
- `body` (string): Response body

**Behavior:**
- GetBody çalıştırır
- Error varsa panic atar
- Aksi takdirde body döner

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Başarılı olursa döner
	body := gosugar.MustGetBody("https://httpbin.org/get")
	fmt.Println("Başarılı:", body[:50])

	// Hata varsa panic
	// body := gosugar.MustGetBody("https://invalid-url")
}
```

---

### 3. `GetJSON[T any](url string) (T, error)`

HTTP GET isteği yapar ve JSON response'u decode eder.

**Signature:**
```go
func GetJSON[T any](url string) (T, error)
```

**Type Parameter:**
- `T`: Decode edilecek struct tipi

**Parametreler:**
- `url` (string): İstek yapılacak URL

**Dönüş Değeri:**
- `result` (T): Decode edilmiş data
- `error`: Hata varsa error, yoksa nil

**Behavior:**
- GetBody çalıştırır
- JSON'u `json.Unmarshal()` ile decode eder
- Başarılı olursa T tipinde değer döner

**Hata Durumları:**
- GetBody hatası: error döner
- JSON decode hatası: error döner

**Örnek:**

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
		fmt.Println("Hata:", err)
		return
	}

	fmt.Printf("Post: %d - %s\n", post.ID, post.Title)
}
```

---

### 4. `GetHeader(url string) (http.Header, error)`

HTTP GET isteği yapar ve response headers'ı döner.

**Signature:**
```go
func GetHeader(url string) (http.Header, error)
```

**Parametreler:**
- `url` (string): İstek yapılacak URL

**Dönüş Değeri:**
- `headers` (http.Header): Response headers
- `error`: Hata varsa error

**Behavior:**
- HTTP GET request yapar
- Status code 200 değilse: error döner
- Headers'ı döner
- `http.Header` case-insensitive map'tir

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	headers, err := gosugar.GetHeader("https://example.com")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Content-Type:", headers.Get("Content-Type"))
	fmt.Println("Server:", headers.Get("Server"))
}
```

---

### 5. `MustGetHeader(url string) http.Header`

`GetHeader` gibi ama hata varsa panic atar.

**Signature:**
```go
func MustGetHeader(url string) http.Header
```

**Parametreler:**
- `url` (string): İstek yapılacak URL

**Dönüş Değeri:**
- `headers` (http.Header): Response headers

**Örnek:**

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

HTTP POST isteği yapar ve body'sini string olarak döner.

**Signature:**
```go
func PostBody(url string, body io.Reader, contentType string) (string, error)
```

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `body` (io.Reader): POST request body
- `contentType` (string): Content-Type header değeri (örn: "application/json")

**Dönüş Değeri:**
- `body` (string): Response body
- `error`: Hata varsa error, yoksa nil

**Behavior:**
- HTTP POST request yapar
- Status code 200 OK değilse: error döner
- Body'sini string'e çevirip döner

**Hata Durumları:**
- Network hatası: error döner
- Non-200 status: `fmt.Errorf("status code: %d")`
- Body okuma hatası: error döner

**Örnek:**

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
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Response:", response[:100])
}
```

---

### 7. `MustPostBody(url string, body io.Reader, contentType string) string`

`PostBody` gibi ama hata varsa panic atar.

**Signature:**
```go
func MustPostBody(url string, body io.Reader, contentType string) string
```

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `body` (io.Reader): POST request body
- `contentType` (string): Content-Type header değeri

**Dönüş Değeri:**
- `body` (string): Response body

**Behavior:**
- PostBody çalıştırır
- Error varsa panic atar
- Aksi takdirde response döner

**Örnek:**

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

HTTP POST isteği yapar JSON payload gönderir ve JSON response'u decode eder.

**Signature:**
```go
func PostJSON[T any](url string, payload any) (T, error)
```

**Type Parameter:**
- `T`: Decode edilecek response struct tipi

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `payload` (any): JSON'a encode edilecek data

**Dönüş Değeri:**
- `result` (T): Decode edilmiş response data
- `error`: Hata varsa error, yoksa nil

**Behavior:**
- Payload'ı JSON'a encode eder
- PostBody çalıştırır
- JSON response'u `json.Unmarshal()` ile decode eder
- Başarılı olursa T tipinde değer döner

**Hata Durumları:**
- JSON encode hatası: error döner
- PostBody hatası: error döner
- JSON decode hatası: error döner

**Örnek:**

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
		Title:  "Yeni Post",
		Body:   "Bu yeni bir post'tur",
		UserID: 1,
	}

	response, err := gosugar.PostJSON[CreatePostResponse](
		"https://jsonplaceholder.typicode.com/posts",
		payload,
	)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Printf("Created Post: %d - %s\n", response.ID, response.Title)
}
```

---

### 9. `PostHeader(url string, body io.Reader, contentType string) (http.Header, error)`

HTTP POST isteği yapar ve response headers'ı döner.

**Signature:**
```go
func PostHeader(url string, body io.Reader, contentType string) (http.Header, error)
```

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `body` (io.Reader): POST request body
- `contentType` (string): Content-Type header değeri

**Dönüş Değeri:**
- `headers` (http.Header): Response headers
- `error`: Hata varsa error

**Behavior:**
- HTTP POST request yapar
- Status code 200 değilse: error döner
- Headers'ı döner
- `http.Header` case-insensitive map'tir

**Örnek:**

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
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Content-Type:", headers.Get("Content-Type"))
	fmt.Println("Server:", headers.Get("Server"))
}
```

---

### 10. `MustPostHeader(url string, body io.Reader, contentType string) http.Header`

`PostHeader` gibi ama hata varsa panic atar.

**Signature:**
```go
func MustPostHeader(url string, body io.Reader, contentType string) http.Header
```

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `body` (io.Reader): POST request body
- `contentType` (string): Content-Type header değeri

**Dönüş Değeri:**
- `headers` (http.Header): Response headers

**Örnek:**

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

HTTP PUT isteği yapar ve body'sini string olarak döner.

**Signature:**
```go
func PutBody(url string, body io.Reader, contentType string) (string, error)
```

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `body` (io.Reader): PUT request body
- `contentType` (string): Content-Type header değeri (örn: "application/json")

**Dönüş Değeri:**
- `body` (string): Response body
- `error`: Hata varsa error, yoksa nil

**Behavior:**
- HTTP PUT request yapar
- Status code 200 OK değilse: error döner
- Body'sini string'e çevirip döner

**Hata Durumları:**
- Network hatası: error döner
- Non-200 status: `fmt.Errorf("status code: %d")`
- Body okuma hatası: error döner

**Örnek:**

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	body := bytes.NewReader([]byte("güncellenen data"))
	
	response, err := gosugar.PutBody(
		"https://httpbin.org/put",
		body,
		"application/json",
	)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Response:", response[:100])
}
```

---

### 12. `MustPutBody(url string, body io.Reader, contentType string) string`

`PutBody` gibi ama hata varsa panic atar.

**Signature:**
```go
func MustPutBody(url string, body io.Reader, contentType string) string
```

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `body` (io.Reader): PUT request body
- `contentType` (string): Content-Type header değeri

**Dönüş Değeri:**
- `body` (string): Response body

**Behavior:**
- PutBody çalıştırır
- Error varsa panic atar
- Aksi takdirde response döner

**Örnek:**

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

HTTP PUT isteği yapar JSON payload gönderir ve JSON response'u decode eder.

**Signature:**
```go
func PutJSON[T any](url string, payload any) (T, error)
```

**Type Parameter:**
- `T`: Decode edilecek response struct tipi

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `payload` (any): JSON'a encode edilecek data

**Dönüş Değeri:**
- `result` (T): Decode edilmiş response data
- `error`: Hata varsa error, yoksa nil

**Behavior:**
- Payload'ı JSON'a encode eder
- PutBody çalıştırır
- JSON response'u `json.Unmarshal()` ile decode eder
- Başarılı olursa T tipinde değer döner

**Hata Durumları:**
- JSON encode hatası: error döner
- PutBody hatası: error döner
- JSON decode hatası: error döner

**Örnek:**

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
		Title: "Güncellenmiş Post",
		Body:  "Bu post güncellenmiştir",
	}

	response, err := gosugar.PutJSON[UpdatePostResponse](
		"https://jsonplaceholder.typicode.com/posts/1",
		payload,
	)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Printf("Güncellenmiş Post: %d - %s\n", response.ID, response.Title)
}
```

---

### 14. `PutHeader(url string, body io.Reader, contentType string) (http.Header, error)`

HTTP PUT isteği yapar ve response headers'ı döner.

**Signature:**
```go
func PutHeader(url string, body io.Reader, contentType string) (http.Header, error)
```

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `body` (io.Reader): PUT request body
- `contentType` (string): Content-Type header değeri

**Dönüş Değeri:**
- `headers` (http.Header): Response headers
- `error`: Hata varsa error

**Behavior:**
- HTTP PUT request yapar
- Status code 200 değilse: error döner
- Headers'ı döner
- `http.Header` case-insensitive map'tir

**Örnek:**

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
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Content-Type:", headers.Get("Content-Type"))
	fmt.Println("Server:", headers.Get("Server"))
}
```

---

### 15. `MustPutHeader(url string, body io.Reader, contentType string) http.Header`

`PutHeader` gibi ama hata varsa panic atar.

**Signature:**
```go
func MustPutHeader(url string, body io.Reader, contentType string) http.Header
```

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `body` (io.Reader): PUT request body
- `contentType` (string): Content-Type header değeri

**Dönüş Değeri:**
- `headers` (http.Header): Response headers

**Örnek:**

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

HTTP DELETE isteği yapar ve body'sini string olarak döner.

**Signature:**
```go
func DeleteBody(url string, body io.Reader, contentType string) (string, error)
```

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `body` (io.Reader): DELETE request body (opsiyonel)
- `contentType` (string): Content-Type header değeri

**Dönüş Değeri:**
- `body` (string): Response body
- `error`: Hata varsa error, yoksa nil

**Behavior:**
- HTTP DELETE request yapar
- Status code 200 OK değilse: error döner
- Body'sini string'e çevirip döner

**Hata Durumları:**
- Network hatası: error döner
- Non-200 status: `fmt.Errorf("status code: %d")`
- Body okuma hatası: error döner

**Örnek:**

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
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Response:", response[:100])
}
```

---

### 17. `MustDeleteBody(url string, body io.Reader, contentType string) string`

`DeleteBody` gibi ama hata varsa panic atar.

**Signature:**
```go
func MustDeleteBody(url string, body io.Reader, contentType string) string
```

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `body` (io.Reader): DELETE request body (opsiyonel)
- `contentType` (string): Content-Type header değeri

**Dönüş Değeri:**
- `body` (string): Response body

**Behavior:**
- DeleteBody çalıştırır
- Error varsa panic atar
- Aksi takdirde response döner

**Örnek:**

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

HTTP DELETE isteği yapar JSON payload gönderir ve JSON response'u decode eder.

**Signature:**
```go
func DeleteJSON[T any](url string, payload any) (T, error)
```

**Type Parameter:**
- `T`: Decode edilecek response struct tipi

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `payload` (any): JSON'a encode edilecek data

**Dönüş Değeri:**
- `result` (T): Decode edilmiş response data
- `error`: Hata varsa error, yoksa nil

**Behavior:**
- Payload'ı JSON'a encode eder
- DeleteBody çalıştırır
- JSON response'u `json.Unmarshal()` ile decode eder
- Başarılı olursa T tipinde değer döner

**Hata Durumları:**
- JSON encode hatası: error döner
- DeleteBody hatası: error döner
- JSON decode hatası: error döner

**Örnek:**

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
		"reason": "artık gerekli değil",
	}

	response, err := gosugar.DeleteJSON[DeleteResponse](
		"https://jsonplaceholder.typicode.com/posts/1",
		payload,
	)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Printf("Silindi: %v\n", response.Deleted)
}
```

---

### 19. `DeleteHeader(url string, body io.Reader, contentType string) (http.Header, error)`

HTTP DELETE isteği yapar ve response headers'ı döner.

**Signature:**
```go
func DeleteHeader(url string, body io.Reader, contentType string) (http.Header, error)
```

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `body` (io.Reader): DELETE request body (opsiyonel)
- `contentType` (string): Content-Type header değeri

**Dönüş Değeri:**
- `headers` (http.Header): Response headers
- `error`: Hata varsa error

**Behavior:**
- HTTP DELETE request yapar
- Status code 200 değilse: error döner
- Headers'ı döner
- `http.Header` case-insensitive map'tir

**Örnek:**

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
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Content-Type:", headers.Get("Content-Type"))
	fmt.Println("Server:", headers.Get("Server"))
}
```

---

### 20. `MustDeleteHeader(url string, body io.Reader, contentType string) http.Header`

`DeleteHeader` gibi ama hata varsa panic atar.

**Signature:**
```go
func MustDeleteHeader(url string, body io.Reader, contentType string) http.Header
```

**Parametreler:**
- `url` (string): İstek yapılacak URL
- `body` (io.Reader): DELETE request body (opsiyonel)
- `contentType` (string): Content-Type header değeri

**Dönüş Değeri:**
- `headers` (http.Header): Response headers

**Örnek:**

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

## Örnekler

### Örnek 1: API Çağrısı

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

	// JSONPlaceholder test API'sinden kullanıcı al
	user, err := gosugar.GetJSON[User](
		"https://jsonplaceholder.typicode.com/users/1",
	)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Printf("Kullanıcı: %s (%s)\n", user.Name, user.Email)
}
```

### Örnek 2: Plain Text Response

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Plain text response al
	body, err := gosugar.GetBody("https://httpbin.org/html")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Response:")
	fmt.Println(body[:200])
}
```

### Örnek 3: HTTP Headers

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	headers, err := gosugar.GetHeader("https://github.com")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("=== GitHub Headers ===")
	fmt.Println("Server:", headers.Get("Server"))
	fmt.Println("Content-Type:", headers.Get("Content-Type"))
	fmt.Println("Connection:", headers.Get("Connection"))
}
```

---

## Limitasyonlar

⚠️ **Şu anki sürüme Kısıtlamalar:**

1. **Sadece 200 OK**: Diğer success status kodları (2xx, 3xx) error sayılır
2. **Custom headers yok**: Authorization vb. headers eklenemez
3. **Timeout yok**: Uzun bağlantılarda bekleyebilir
4. **Redirect handling yok**: Yönlendirmeleri takip etmez
5. **Request iptal edilemez**: Başlanmış istekler iptal edilemez

**Workaround:** `net/http` package'ını direkt kullanın.

---

## İlişkili Modüller

- **`errors.go`**: Error handling
- **`getting-started.md`**: Başlama rehberi

