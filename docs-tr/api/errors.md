# API Referansı: errors - Hata Yönetimi

Panic handling ve error recovery mekanizmaları sağlayan modül. Güvenli error handling patterns sunun.

## 📋 İçindekiler

- [Genel Bakış](#genel-bakış)
- [Fonksiyonlar](#fonksiyonlar)
- [Örnekler](#örnekler)
- [Patterns](#patterns)

---

## Genel Bakış

### Amaç

- Panic ile error handling
- Try/catch benzeri güvenli execution
- Fallback değerler sağlamak

### Başlıca Özellikler

- ✅ Generic type support
- ✅ Panic-safe execution (`Try`)
- ✅ Fallback mekanizması (`Or`)
- ✅ Must pattern

---

## Fonksiyonlar

### 1. `Must[T any](v T, err error) T`

`(T, error)` dönen fonksiyonlardan error varsa panic atar.

**Signature:**
```go
func Must[T any](v T, err error) T
```

**Type Parameter:**
- `T`: Dönüş tipi

**Parametreler:**
- `v` (T): Başarılı durumda dönülecek değer
- `err` (error): Hata (nil ise başarılı)

**Dönüş Değeri:**
- Başarılı ise: `v`
- Hata ise: panic atar

**Behavior:**
- `err != nil` ise panic atar
- Aksi takdirde `v` döner

**Örnek:**

```go
package main

import (
	"os"
	"github.com/coderianx/gosugar"
)

func main() {
	// Başarılı (hata yok)
	file := gosugar.Must(os.Open("data.txt"))
	defer file.Close()

	// Hata varsa panic
	// gosugar.Must(os.Open("nonexistent.txt")) // panic!
}
```

---

### 2. `Check(err error)`

Yalnızca error dönen fonksiyonlar için. Error varsa panic atar.

**Signature:**
```go
func Check(err error)
```

**Parametreler:**
- `err` (error): Kontrol edilecek hata

**Behavior:**
- `err != nil` ise panic atar

**Örnek:**

```go
package main

import (
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Başarılı
	gosugar.Check(os.Mkdir("./data", 0755))

	// Hata varsa panic
	// gosugar.Check(os.RemoveAll("/")) // panic!
}
```

---

### 3. `Try[T any](fn func() T) (T, bool)`

Panic'ten kurtulur. Panic atan kodu güvenli çalıştırır.

**Signature:**
```go
func Try[T any](fn func() T) (T, bool)
```

**Type Parameter:**
- `T`: Fonksiyonun dönüş tipi

**Parametreler:**
- `fn` (func() T): Çalıştırılacak fonksiyon

**Dönüş Değeri:**
- `v` (T): Fonksiyonun dönüş değeri (panic ise zero-value)
- `ok` (bool): Başarılı ise `true`, panic ise `false`

**Behavior:**
- `fn()` çalıştırılır
- Panic varsa recover eder ve `ok=false` döner
- Aksi takdirde `ok=true` döner

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Riskli kod
	value, ok := gosugar.Try(func() int {
		// Bu panic atabilir
		return 100 / 0
	})

	if !ok {
		fmt.Println("Hata: Kod panic attı")
	} else {
		fmt.Println("Başarılı:", value)
	}
}
```

---

### 4. `Or[T any](v T, ok bool, fallback T) T`

Try ile birlikte kullanılır. Fallback değer sağlar.

**Signature:**
```go
func Or[T any](v T, ok bool, fallback T) T
```

**Type Parameters:**
- `T`: Değer tipi

**Parametreler:**
- `v` (T): Ana değer
- `ok` (bool): Başarılı olup olmadığı (`Try` dönüş değeri)
- `fallback` (T): Başarısız ise kullanılacak değer

**Dönüş Değeri:**
- `ok=true` ise: `v`
- `ok=false` ise: `fallback`

**Behavior:**
- Simple ternary operator gibi

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	value, ok := gosugar.Try(func() int {
		return 100 / 0 // panic
	})

	result := gosugar.Or(value, ok, 0)
	fmt.Println("Sonuç:", result) // 0 (fallback)
}
```

---

### 5. `Ignore(err error)`

Error'u bilinçli şekilde yutmak için.

**Signature:**
```go
func Ignore(err error)
```

**Parametreler:**
- `err` (error): Yutulacak error

**Behavior:**
- Error'u göz ardı eder
- Linter uyarılarını kaldırmak için faydalı

**Örnek:**

```go
package main

import (
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Hata olsa da umursamıyoruz
	gosugar.Ignore(os.RemoveAll("./temp"))
}
```

---

## Örnekler

### Örnek 1: Must Pattern

```go
package main

import (
	"os"
	"github.com/coderianx/gosugar"
)

func main() {
	// Dosya açma
	file := gosugar.Must(os.Open("config.json"))
	defer file.Close()

	// Başarılı ise devam et
	println("Dosya açıldı")
}
```

### Örnek 2: Try/Or Pattern

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"strconv"
)

func main() {
	// Risky: string'i integer'a çevir
	value, ok := gosugar.Try(func() int {
		return gosugar.Must(strconv.Atoi("abc"))
	})

	// Başarısız olursa 0 kullan
	result := gosugar.Or(value, ok, 0)
	fmt.Println("Değer:", result)
}
```

### Örnek 3: File Operations

```go
package main

import (
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Başarısız olabilecek operasyonlar
	content, ok := gosugar.Try(func() string {
		return gosugar.ReadFile("data.txt")
	})

	if ok {
		println("Okunan:", content)
	} else {
		println("Dosya okunamadı, varsayılan kullanılıyor")
		content = "Varsayılan veri"
	}
}
```

### Örnek 4: Custom Function with Try

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func divideByZero() int {
	return 100 / 0 // panic!
}

func main() {
	result, ok := gosugar.Try(divideByZero)
	if !ok {
		fmt.Println("İşlem başarısız")
		result = 0
	}
	fmt.Println("Sonuç:", result)
}
```

---

## Patterns

### Pattern 1: Startup Validation

Başlangıçta zorunlu değişkenleri kontrol et:

```go
func main() {
	// Hata varsa panic (başlangıçta catch edilebilir)
	port := gosugar.Must(strconv.Atoi(os.Getenv("PORT")))
	dbURL := gosugar.MustEnv("DATABASE_URL")
	
	// Başarılı olursa devam
	println("Port:", port)
}
```

### Pattern 2: Fallback Values

Opsiyonel işlemler için:

```go
func main() {
	content, ok := gosugar.Try(func() string {
		return gosugar.ReadFile("config.json")
	})
	
	config := gosugar.Or(content, ok, "{}")
}
```

### Pattern 3: Multiple Try/Or

```go
func main() {
	// Sıra ile dene
	v1, ok1 := gosugar.Try(func() string {
		return gosugar.ReadFile("config.local.json")
	})
	
	v2, ok2 := gosugar.Try(func() string {
		return gosugar.ReadFile("config.json")
	})
	
	config := gosugar.Or(v1, ok1, gosugar.Or(v2, ok2, "{}"))
}
```

---

## İlişkili Modüller

- **`env.go`**: Ortam değişkenleri (MustEnv)
- **`file.go`**: Dosya işlemleri (hata handling)
- **`getting-started.md`**: Başlama rehberi

