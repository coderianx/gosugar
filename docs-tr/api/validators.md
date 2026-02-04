# API Referansı: validators - Input Doğrulama

Kullanıcı girdilerini kontrol etmek için hazır ve composable validatörler sağlayan modül.

## 📋 İçindekiler

- [Genel Bakış](#genel-bakış)
- [Validatör Türü](#validatör-türü)
- [Hazır Validatörler](#hazır-validatörler)
- [Özel Validatör Yazma](#özel-validatör-yazma)
- [Örnekler](#örnekler)

---

## Genel Bakış

### Amaç

- Kullanıcı girdilerini doğrulamak
- Zincirleme şekilde validatörler uygulamak
- Özel validatörler yazabilir yapmak

### Başlıca Özellikler

- ✅ Composable validatörler
- ✅ Funktif programming pattern
- ✅ Hata mesajları
- ✅ Genişletilebilir tasarım

---

## Validatör Türü

### `Validator`

```go
type Validator func(string) error
```

**Açıklama:**
- Bir string parameter alır
- Validasyon başarılı olursa `nil` döner
- Validasyon başarısız olursa `error` döner

**Kullanım:**

```go
// Hazır validatör
notEmpty := gosugar.NotEmpty()
err := notEmpty("")           // error döner
err = notEmpty("hello")       // nil döner

// Validatörü Input'a geç
email := gosugar.Input(
	"E-mail: ",
	gosugar.NotEmpty(),  // Burası Validator fonksiyon
)
```

---

## Hazır Validatörler

### 1. `NotEmpty() Validator`

Boş string'i redder.

**Signature:**
```go
func NotEmpty() Validator
```

**Behavior:**
- Input boş string ise: error döner ("value cannot be empty")
- Input boş değilse: nil döner

**Örnek:**

```go
username := gosugar.Input(
	"Kullanıcı adı: ",
	gosugar.NotEmpty(),
)
// Boş giriş varsa "invalid string input" hatası
```

---

### 2. `MinLen(n int) Validator`

Minimum karakter sayısı kontrolü yapar.

**Signature:**
```go
func MinLen(n int) Validator
```

**Parametreler:**
- `n` (int): Minimum karakter sayısı

**Behavior:**
- `len(string) < n` ise: error döner
- Örneğin: `len("hi") < 3` → error

**Error Mesajı:**
```
"minimum length is 3"
```

**Örnek:**

```go
password := gosugar.Input(
	"Şifre (min 8): ",
	gosugar.NotEmpty(),
	gosugar.MinLen(8),
)
// "1234" girişi: MinLen(8) hatası verir
```

---

### 3. `MaxLen(n int) Validator`

Maksimum karakter sayısı kontrolü yapar.

**Signature:**
```go
func MaxLen(n int) Validator
```

**Parametreler:**
- `n` (int): Maksimum karakter sayısı

**Behavior:**
- `len(string) > n` ise: error döner

**Error Mesajı:**
```
"maximum length is 100"
```

**Örnek:**

```go
bio := gosugar.Input(
	"Biyografi (max 200): ",
	gosugar.MaxLen(200),
)
```

---

## Özel Validatör Yazma

Validatörler fonksiyon olduğu için, kendi validatörleriniz yazabilirsiniz:

### Pattern 1: Basit Validatör

```go
package main

import "github.com/coderianx/gosugar"

// Sadece sayılar içeren string
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
		"Telefon: ",
		NumericOnly(),
		gosugar.MinLen(10),
	)
	println(phoneNumber)
}
```

### Pattern 2: Regex Validatör

```go
package main

import (
	"fmt"
	"regexp"
	"github.com/coderianx/gosugar"
)

// E-mail formatı kontrolü
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
		"E-mail: ",
		gosugar.NotEmpty(),
		EmailFormat(),
	)
	println(email)
}
```

### Pattern 3: Parametreli Validatör

```go
// "A", "B", "C" içinden seçim
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
		"Seviye (LOW/MEDIUM/HIGH): ",
		OneOf("LOW", "MEDIUM", "HIGH"),
	)
	println(level)
}
```

---

## Örnekler

### Örnek 1: Kombinasyon

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	username := gosugar.Input(
		"Kullanıcı adı: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(3),
		gosugar.MaxLen(20),
	)
	println("Kullanıcı adı:", username)
}
```

### Örnek 2: Farklı Validatörler

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Başlık: 5-100 karakter
	title := gosugar.Input(
		"Başlık: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(5),
		gosugar.MaxLen(100),
	)

	// Açıklama: 20-1000 karakter
	description := gosugar.Input(
		"Açıklama: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(20),
		gosugar.MaxLen(1000),
	)

	fmt.Println("Kaydedildi")
}
```

### Örnek 3: Custom + Hazır Validatörler

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"strings"
)

// Sadece harfler
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
		"İsim: ",
		gosugar.NotEmpty(),
		LettersOnly(),
		gosugar.MinLen(2),
		gosugar.MaxLen(50),
	)

	lastName := gosugar.Input(
		"Soyadı: ",
		gosugar.NotEmpty(),
		LettersOnly(),
		gosugar.MinLen(2),
		gosugar.MaxLen(50),
	)

	fmt.Printf("Hoş geldiniz, %s %s!\n", firstName, lastName)
}
```

---

## İlişkili Modüller

- **`input.go`**: Validatörler ile input alma
- **`errors.go`**: Error handling
- **`design-patterns.md`**: Özel validatör yazma örnekleri

