# API Referansı: random - Rastgele Veri Üretimi

Rastgele sayılar, stringler ve seçimler üreten modül. Test ve demo için kullanışlı.

## 📋 İçindekiler

- [Genel Bakış](#genel-bakış)
- [Fonksiyonlar](#fonksiyonlar)
- [Örnekler](#örnekler)

---

## Genel Bakış

### Amaç

- Rastgele integer, float, boolean üretmek
- Rastgele string oluşturmak
- Listeden rastgele element seçmek

### Başlıca Özellikler

- ✅ Seed otomatik başlatılır
- ✅ Type-safe generics
- ✅ Farklı aralıklar (inclusive/exclusive)
- ✅ Hata validasyonu

---

## Fonksiyonlar

### 1. `RandInt(min, max int) int`

Belirtilen aralıkta rastgele integer döner.

**Signature:**
```go
func RandInt(min, max int) int
```

**Parametreler:**
- `min` (int): Minimum değer (dahil)
- `max` (int): Maksimum değer (dahil)

**Dönüş Değeri:**
- Rastgele integer: `min <= x <= max`

**Behavior:**
- `min > max` ise panic atar
- Her çağrıda farklı rastgele sayı üretir

**Hata Durumları:**
- `min > max`: `panic("min cannot be greater than max")`

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Zar atma (1-6)
	dice := gosugar.RandInt(1, 6)
	fmt.Println("Zar:", dice)

	// 1-100 arası
	num := gosugar.RandInt(1, 100)
	fmt.Println("Rastgele:", num)

	// Negatif sayılar
	val := gosugar.RandInt(-10, 10)
	fmt.Println("Değer:", val)
}
```

---

### 2. `RandFloat(min, max float64) float64`

Belirtilen aralıkta rastgele float döner.

**Signature:**
```go
func RandFloat(min, max float64) float64
```

**Parametreler:**
- `min` (float64): Minimum değer (dahil)
- `max` (float64): Maksimum değer (hariç)

**Dönüş Değeri:**
- Rastgele float64: `min <= x < max`

**Behavior:**
- `min >= max` ise panic atar
- **Maksimum hariçtir** (0.0-1.0 aralığı 1.0 içermez)

**Hata Durumları:**
- `min >= max`: `panic("min must be less than max")`

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// 0.0-1.0 arası (olasılık)
	chance := gosugar.RandFloat(0.0, 1.0)
	fmt.Printf("Şans: %.4f\n", chance)

	// 10.5-20.5 arası
	price := gosugar.RandFloat(10.5, 20.5)
	fmt.Printf("Fiyat: $%.2f\n", price)
}
```

---

### 3. `RandBool() bool`

Rastgele boolean değer döner.

**Signature:**
```go
func RandBool() bool
```

**Dönüş Değeri:**
- `true` veya `false` (50/50 şans)

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	if gosugar.RandBool() {
		fmt.Println("Yazı")
	} else {
		fmt.Println("Tura")
	}
}
```

---

### 4. `RandString(length int) string`

Belirtilen uzunlukta rastgele string döner (sadece harfler).

**Signature:**
```go
func RandString(length int) string
```

**Parametreler:**
- `length` (int): String uzunluğu

**Dönüş Değeri:**
- Rastgele string (A-Z ve a-z karakterleri)

**Behavior:**
- Sadece İngilizce harfler (26 + 26 = 52 karakter)
- `length <= 0` ise panic atar

**Hata Durumları:**
- `length <= 0`: `panic("length must be positive")`

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Token oluştur (10 karakter)
	token := gosugar.RandString(10)
	fmt.Println("Token:", token)

	// ID oluştur (32 karakter)
	id := gosugar.RandString(32)
	fmt.Println("ID:", id)
}
```

---

### 5. `Choice[T any](items []T) T`

Listeden rastgele bir element seçer.

**Signature:**
```go
func Choice[T any](items []T) T
```

**Type Parameter:**
- `T`: Herhangi bir tür (generic)

**Parametreler:**
- `items` ([]T): Listeden seçim yapılacak

**Dönüş Değeri:**
- Rastgele seçilmiş element

**Behavior:**
- Liste boş ise panic atar
- Type-safe (compile-time kontrol)

**Hata Durumları:**
- Boş liste: `panic("cannot choose from empty slice")`

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// String'lerden seçim
	colors := []string{"red", "green", "blue", "yellow"}
	color := gosugar.Choice(colors)
	fmt.Println("Renk:", color)

	// Integer'lardan seçim
	numbers := []int{10, 20, 30, 40, 50}
	num := gosugar.Choice(numbers)
	fmt.Println("Sayı:", num)

	// Struct'lardan seçim
	type User struct {
		Name string
	}
	users := []User{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Charlie"},
	}
	selected := gosugar.Choice(users)
	fmt.Println("Seçilen:", selected.Name)
}
```

---

## Örnekler

### Örnek 1: Oyun

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	fmt.Println("🎮 Zar Oyunu")
	fmt.Println("3 zarı at...\n")

	total := 0
	for i := 1; i <= 3; i++ {
		dice := gosugar.RandInt(1, 6)
		fmt.Printf("Zar %d: %d\n", i, dice)
		total += dice
	}

	fmt.Printf("\nToplam: %d\n", total)
}
```

### Örnek 2: Rastgele Seçim

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Rastgele gün seç
	days := []string{"Pazartesi", "Salı", "Çarşamba", "Perşembe", "Cuma", "Cumartesi", "Pazar"}
	day := gosugar.Choice(days)
	fmt.Println("Rastgele gün:", day)

	// Rastgele öncelik seç
	priorities := []string{"LOW", "MEDIUM", "HIGH"}
	priority := gosugar.Choice(priorities)
	fmt.Println("Öncelik:", priority)
}
```

### Örnek 3: Rastgele Token/ID

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

	// Şifre reset kodu
	code := gosugar.RandString(6)
	fmt.Println("Kod:", code)

	// Session ID
	sessionID := gosugar.RandString(64)
	fmt.Println("Session:", sessionID)
}
```

### Örnek 4: Test Verisi

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	fmt.Println("=== Test Verisi Üretimi ===\n")

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

## İlişkili Modüller

- **`errors.go`**: Error handling
- **`getting-started.md`**: İlk adımlar

