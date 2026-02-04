# API Referansı: input - Kullanıcıdan Input Alma

Terminal üzerinden kullanıcıdan veri almayı sağlayan modül. String, integer, float değerleri interaktif şekilde alabilirsiniz.

## 📋 İçindekiler

- [Genel Bakış](#genel-bakış)
- [Fonksiyonlar](#fonksiyonlar)
- [Örnekler](#örnekler)

---

## Genel Bakış

### Amaç

- Kullanıcıdan terminal üzerinden veri almak
- Girdileri validatörler ile kontrol etmek
- Geçersiz girdide varsayılan değer döndürmek

### Başlıca Özellikler

- ✅ String, Integer, Float input
- ✅ Composable validatörler
- ✅ Varsayılan değer desteği
- ✅ Otomatik whitespace kaldırma
- ✅ Panic-based error handling

---

## Fonksiyonlar

### 1. `Input(prompt string, validators ...Validator) string`

Kullanıcıdan string input alır ve validatörlerle kontrol eder.

**Signature:**
```go
func Input(prompt string, validators ...Validator) string
```

**Parametreler:**
- `prompt` (string): Gösterilecek soru/rehber metni
- `validators` (variadic): Uygulanacak validatörler (opsiyonel)

**Dönüş Değeri:**
- Kullanıcının girdiği string (boşluklar kaldırılmış)

**Behavior:**
- Prompt'u gösterir ve girdiye bekler
- Girdiyi `strings.TrimSpace()` ile temizler
- Her validatörü sıra ile çalıştırır
- Validasyon başarısız olursa panic atar
- Validasyon başarılı olursa değeri döner

**Hata Durumları:**
- Validasyon hatası: `panic("invalid string input: ...")`
- Input okuma hatası: `panic("input error")`

**Örnek:**

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	// Basit input (validatör yok)
	name := gosugar.Input("Adınız: ")
	println("Merhaba,", name)

	// Validatörlerle input
	email := gosugar.Input(
		"E-mail: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(5),
	)
	println("E-mail:", email)
}
```

**Execution:**
```
Adınız: John Doe
Merhaba, John Doe
E-mail: ab@test.com   # MinLen(5) hatası! Tekrar sor
E-mail: valid@email.com
E-mail: valid@email.com
```

---

### 2. `InputInt(prompt string, defaultValue ...int) int`

Kullanıcıdan integer input alır. Geçersiz input varsa varsayılan değer döner.

**Signature:**
```go
func InputInt(prompt string, defaultValue ...int) int
```

**Parametreler:**
- `prompt` (string): Gösterilecek soru
- `defaultValue` (variadic): Geçersiz girdide dönecek varsayılan değer

**Dönüş Değeri:**
- Integer değer (başarılı ise) veya varsayılan (başarısız ise)

**Behavior:**
- Prompt'u gösterir
- `strconv.Atoi()` ile integer'a çevirmeye çalışır
- Başarılı olursa: integer döner
- Başarısız olursa:
  - Varsayılan değer varsa: onu döner
  - Yoksa: panic atar

**Hata Durumları:**
- Geçersiz format ve varsayılan yok: `panic("invalid integer input: ...")`

**Örnek:**

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	// Başarılı input
	age := gosugar.InputInt("Yaşınız: ")
	println("Yaş:", age)

	// Geçersiz input, varsayılan döner
	port := gosugar.InputInt("Port (varsayılan 8080): ", 8080)
	println("Port:", port)
}
```

**Execution:**
```
Yaşınız: abc       # Geçersiz, panic!
Port (varsayılan 8080): xyz
Port: 8080         # Varsayılan döndü, hata yok
```

---

### 3. `InputFloat(prompt string, defaultValue ...float64) float64`

Kullanıcıdan float input alır. Geçersiz input varsa varsayılan değer döner.

**Signature:**
```go
func InputFloat(prompt string, defaultValue ...float64) float64
```

**Parametreler:**
- `prompt` (string): Gösterilecek soru
- `defaultValue` (variadic): Geçersiz girdide varsayılan

**Dönüş Değeri:**
- Float64 değer

**Behavior:**
- Prompt'u gösterir
- `strconv.ParseFloat()` ile dönüştürmeye çalışır
- Başarılı: float döner
- Başarısız: varsayılan döner veya panic

**Örnek:**

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	price := gosugar.InputFloat("Fiyat: ", 9.99)
	println("Fiyat:", price)

	discount := gosugar.InputFloat("İndirim oranı (0-1): ")
	println("İndirim:", discount)
}
```

---

## Örnekler

### Örnek 1: Basit Anket

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	fmt.Println("=== Anket Formu ===\n")

	name := gosugar.Input(
		"Adınız: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(2),
	)

	age := gosugar.InputInt("Yaşınız: ", 0)

	email := gosugar.Input(
		"E-mail: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(5),
	)

	fmt.Printf("\nTeşekkürler %s! Bilgileriniz kaydedildi.\n", name)
}
```

### Örnek 2: Validatörler İle

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Kullanıcı adı: 3-20 karakter
	username := gosugar.Input(
		"Kullanıcı adı (3-20 karakter): ",
		gosugar.NotEmpty(),
		gosugar.MinLen(3),
		gosugar.MaxLen(20),
	)

	// Şifre: minimum 8 karakter
	password := gosugar.Input(
		"Şifre (min 8 karakter): ",
		gosugar.NotEmpty(),
		gosugar.MinLen(8),
	)

	fmt.Printf("Kayıt başarılı: %s\n", username)
}
```

### Örnek 3: Sayısal Input

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	quantity := gosugar.InputInt(
		"Miktar: ",
		1,  // varsayılan: 1
	)

	price := gosugar.InputFloat(
		"Birim Fiyat (₺): ",
		0.0,  // varsayılan: 0
	)

	total := float64(quantity) * price
	fmt.Printf("Toplam: ₺%.2f\n", total)
}
```

---

## İlişkili Modüller

- **`validators.go`**: Validatör türleri ve hazır validatörler
- **`errors.go`**: Error handling
- **`env.go`**: Çevre değişkenleriyle default sağlama

