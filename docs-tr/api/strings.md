# API Referansı: strings - String Manipülasyon Yardımcıları

Yaygın string işlemleri, case dönüşümleri ve metin formatlama için yardımcı fonksiyonlar sağlayan kapsamlı modül.

## 📋 İçindekiler

- [Genel Bakış](#genel-bakış)
- [Fonksiyonlar](#fonksiyonlar)
- [Örnekler](#örnekler)

---

## Genel Bakış

### Amaç

- Yaygın string işlemlerini basitleştirmek
- Case dönüşümü yardımcıları sağlamak
- String padding ve truncation işlemleri
- URL-uyumlu slug oluşturmak
- Tekrarlanan string manipülasyon kodunu azaltmak

### Başlıca Özellikler

- ✅ UTF-8 güvenli işlemler
- ✅ Case dönüşümleri (camelCase, snake_case, kebab-case, PascalCase)
- ✅ Padding ve truncation
- ✅ String trim ve reverse
- ✅ Slug üretimi
- ✅ Standart kütüphane fonksiyonları için kolay wrapperlar

---

## Fonksiyonlar

### 1. `Trim(s string) string`

String'in başından ve sonundan boşluk karakterlerini kaldırır.

**Signature:**
```go
func Trim(s string) string
```

**Parametreler:**
- `s` (string): Kırpılacak string

**Dönüş Değeri:**
- `string`: Kırpılmış string

**Behavior:**
- Tüm başındaki ve sondaki boşlukları kaldırır
- İç boşlukları korur
- Boş string giriş için boş string döner

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Trim("  merhaba dünya  ")
	fmt.Println(result) // "merhaba dünya"
	
	result = gosugar.Trim("\t\n  metin  \n\t")
	fmt.Println(result) // "metin"
}
```

---

### 2. `ToUpper(s string) string`

String'i büyük harflere dönüştürür.

**Signature:**
```go
func ToUpper(s string) string
```

**Parametreler:**
- `s` (string): Dönüştürülecek string

**Dönüş Değeri:**
- `string`: Büyük harfli string

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.ToUpper("merhaba")
	fmt.Println(result) // "MERHABA"
	
	result = gosugar.ToUpper("GoSugar")
	fmt.Println(result) // "GOSUGAR"
}
```

---

### 3. `ToLower(s string) string`

String'i küçük harflere dönüştürür.

**Signature:**
```go
func ToLower(s string) string
```

**Parametreler:**
- `s` (string): Dönüştürülecek string

**Dönüş Değeri:**
- `string`: Küçük harfli string

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.ToLower("MERHABA")
	fmt.Println(result) // "merhaba"
	
	result = gosugar.ToLower("GoSugar")
	fmt.Println(result) // "gosugar"
}
```

---

### 4. `Reverse(s string) string`

String'i karakter karakter tersine çevirir (UTF-8 güvenli).

**Signature:**
```go
func Reverse(s string) string
```

**Parametreler:**
- `s` (string): Tersine çevrilecek string

**Dönüş Değeri:**
- `string`: Ters çevrilmiş string

**Behavior:**
- Multi-byte UTF-8 karakterleri doğru şekilde işler
- Emoji ve özel karakterleri düzgün tersine çevirir

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Reverse("merhaba")
	fmt.Println(result) // "abahrem"
	
	result = gosugar.Reverse("😀🎉")
	fmt.Println(result) // "🎉😀"
}
```

---

### 5. `Contains(s, substr string) bool`

String'in bir alt string içerip içermediğini kontrol eder.

**Signature:**
```go
func Contains(s, substr string) bool
```

**Parametreler:**
- `s` (string): Aranacak string
- `substr` (string): Aranan alt string

**Dönüş Değeri:**
- `bool`: Alt string bulunursa true, değilse false

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	found := gosugar.Contains("merhaba dünya", "dünya")
	fmt.Println(found) // true
	
	found = gosugar.Contains("merhaba", "xyz")
	fmt.Println(found) // false
}
```

---

### 6. `HasPrefix(s, prefix string) bool`

String'in belirtilen önek ile başlayıp başlamadığını kontrol eder.

**Signature:**
```go
func HasPrefix(s, prefix string) bool
```

**Parametreler:**
- `s` (string): Kontrol edilecek string
- `prefix` (string): Aranacak önek

**Dönüş Değeri:**
- `bool`: Önek varsa true, yoksa false

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	has := gosugar.HasPrefix("merhaba.go", "merhaba")
	fmt.Println(has) // true
	
	has = gosugar.HasPrefix("merhaba.go", ".go")
	fmt.Println(has) // false
}
```

---

### 7. `HasSuffix(s, suffix string) bool`

String'in belirtilen sonek ile bitip bitmediğini kontrol eder.

**Signature:**
```go
func HasSuffix(s, suffix string) bool
```

**Parametreler:**
- `s` (string): Kontrol edilecek string
- `suffix` (string): Aranacak sonek

**Dönüş Değeri:**
- `bool`: Sonek varsa true, yoksa false

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	has := gosugar.HasSuffix("merhaba.go", ".go")
	fmt.Println(has) // true
	
	has = gosugar.HasSuffix("merhaba", ".go")
	fmt.Println(has) // false
}
```

---

### 8. `Replace(s, old, new string) string`

İlk oluşum olan alt string'i değiştirir.

**Signature:**
```go
func Replace(s, old, new string) string
```

**Parametreler:**
- `s` (string): İşlenecek string
- `old` (string): Bulunacak alt string
- `new` (string): Değiştirme alt string

**Dönüş Değeri:**
- `string`: İlk oluşum değiştirilen string

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Replace("merhaba merhaba", "merhaba", "selam")
	fmt.Println(result) // "selam merhaba"
	
	result = gosugar.Replace("kedi ve kedi", "kedi", "köpek")
	fmt.Println(result) // "köpek ve kedi"
}
```

---

### 9. `ReplaceAll(s, old, new string) string`

Tüm oluşumları olan alt string'i değiştirir.

**Signature:**
```go
func ReplaceAll(s, old, new string) string
```

**Parametreler:**
- `s` (string): İşlenecek string
- `old` (string): Bulunacak alt string
- `new` (string): Değiştirme alt string

**Dönüş Değeri:**
- `string`: Tüm oluşumları değiştirilen string

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.ReplaceAll("merhaba merhaba", "merhaba", "selam")
	fmt.Println(result) // "selam selam"
	
	result = gosugar.ReplaceAll("kedi ve kedi", "kedi", "köpek")
	fmt.Println(result) // "köpek ve köpek"
}
```

---

### 10. `Split(s, sep string) []string`

String'i ayırıcıya göre böler.

**Signature:**
```go
func Split(s, sep string) []string
```

**Parametreler:**
- `s` (string): Bölünecek string
- `sep` (string): Ayırıcı string

**Dönüş Değeri:**
- `[]string`: Alt string'lerin slice'ı

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	parts := gosugar.Split("a,b,c", ",")
	fmt.Println(parts) // ["a" "b" "c"]
	
	words := gosugar.Split("merhaba-dünya-go", "-")
	fmt.Println(words) // ["merhaba" "dünya" "go"]
}
```

---

### 11. `Join(strs []string, sep string) string`

String slice'ını ayırıcı ile birleştirir.

**Signature:**
```go
func Join(strs []string, sep string) string
```

**Parametreler:**
- `strs` ([]string): Birleştirilecek string slice'ı
- `sep` (string): Ayırıcı string

**Dönüş Değeri:**
- `string`: Birleştirilmiş string

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Join([]string{"a", "b", "c"}, ",")
	fmt.Println(result) // "a,b,c"
	
	result = gosugar.Join([]string{"merhaba", "dünya"}, " ")
	fmt.Println(result) // "merhaba dünya"
}
```

---

### 12. `Repeat(s string, count int) string`

String'i belirtilen sayıda tekrarlar.

**Signature:**
```go
func Repeat(s string, count int) string
```

**Parametreler:**
- `s` (string): Tekrarlanacak string
- `count` (int): Tekrarlama sayısı

**Dönüş Değeri:**
- `string`: Tekrarlanmış string

**Panics:**
- count negatifse

**Örnek:**

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

String'i soldan belirtilen genişliğe kadar doldurur.

**Signature:**
```go
func PadLeft(s string, width int, char rune) string
```

**Parametreler:**
- `s` (string): Doldurulacak string
- `width` (int): Hedef genişlik
- `char` (rune): Doldurma karakteri

**Dönüş Değeri:**
- `string`: Soldan doldurulmuş string

**Behavior:**
- String zaten width'ten geniş ise, değiştirilmeden döner
- Byte sayısını değil, karakter sayısını dikkate alır

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.PadLeft("5", 3, '0')
	fmt.Println(result) // "005"
	
	result = gosugar.PadLeft("merhaba", 10, '-')
	fmt.Println(result) // "---merhaba"
}
```

---

### 14. `PadRight(s string, width int, char rune) string`

String'i sağdan belirtilen genişliğe kadar doldurur.

**Signature:**
```go
func PadRight(s string, width int, char rune) string
```

**Parametreler:**
- `s` (string): Doldurulacak string
- `width` (int): Hedef genişlik
- `char` (rune): Doldurma karakteri

**Dönüş Değeri:**
- `string`: Sağdan doldurulmuş string

**Behavior:**
- String zaten width'ten geniş ise, değiştirilmeden döner
- Byte sayısını değil, karakter sayısını dikkate alır

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.PadRight("5", 3, '0')
	fmt.Println(result) // "500"
	
	result = gosugar.PadRight("merhaba", 10, '-')
	fmt.Println(result) // "merhaba---"
}
```

---

### 15. `Truncate(s string, length int, suffix string) string`

String'i maksimum uzunluğa kırpar ve kesilirse sonek ekler.

**Signature:**
```go
func Truncate(s string, length int, suffix string) string
```

**Parametreler:**
- `s` (string): Kırpılacak string
- `length` (int): Maksimum uzunluk (sonek dahil)
- `suffix` (string): Kırpılırsa eklenecek sonek (genellikle "...")

**Dönüş Değeri:**
- `string`: Sonek ile kesilmiş string

**Behavior:**
- String length'ten kısa veya eşit ise, değiştirilmeden döner
- Sonek, length hesaplamasına dahil edilir

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Truncate("merhaba dünya", 8, "...")
	fmt.Println(result) // "merhaba..."
	
	result = gosugar.Truncate("hi", 8, "...")
	fmt.Println(result) // "hi"
}
```

---

### 16. `Capitalize(s string) string`

String'in ilk karakterini büyük harfe dönüştürür.

**Signature:**
```go
func Capitalize(s string) string
```

**Parametreler:**
- `s` (string): Büyük harfe dönüştürülecek string

**Dönüş Değeri:**
- `string`: İlk karakteri büyük harfli string

**Behavior:**
- Sadece ilk karakteri etkiler
- String'in geri kalanı değiştirilmez
- Giriş boş ise boş string döner

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Capitalize("merhaba")
	fmt.Println(result) // "Merhaba"
	
	result = gosugar.Capitalize("MERHABA")
	fmt.Println(result) // "MERHABA"
}
```

---

### 17. `Decapitalize(s string) string`

String'in ilk karakterini küçük harfe dönüştürür.

**Signature:**
```go
func Decapitalize(s string) string
```

**Parametreler:**
- `s` (string): Küçük harfe dönüştürülecek string

**Dönüş Değeri:**
- `string`: İlk karakteri küçük harfli string

**Behavior:**
- Sadece ilk karakteri etkiler
- String'in geri kalanı değiştirilmez
- Giriş boş ise boş string döner

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Decapitalize("Merhaba")
	fmt.Println(result) // "merhaba"
	
	result = gosugar.Decapitalize("MERHABA")
	fmt.Println(result) // "mERHABA"
}
```

---

### 18. `CamelCase(s string) string`

String'i camelCase'e dönüştürür.

**Signature:**
```go
func CamelCase(s string) string
```

**Parametreler:**
- `s` (string): Dönüştürülecek string

**Dönüş Değeri:**
- `string`: camelCase versiyonu

**Behavior:**
- Boşlukları, kısa çizgileri ve alt çizgileri kaldırır
- İlk sözcük küçük harfliyken, sonraki sözcükler büyük harfliydir
- Birden fazla ayırıcıyı işler

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.CamelCase("merhaba dünya")
	fmt.Println(result) // "merhabaDünya"
	
	result = gosugar.CamelCase("merhaba-dünya")
	fmt.Println(result) // "merhabaDünya"
	
	result = gosugar.CamelCase("merhaba_dünya")
	fmt.Println(result) // "merhabaDünya"
}
```

---

### 19. `SnakeCase(s string) string`

String'i snake_case'e dönüştürür.

**Signature:**
```go
func SnakeCase(s string) string
```

**Parametreler:**
- `s` (string): Dönüştürülecek string

**Dönüş Değeri:**
- `string`: snake_case versiyonu

**Behavior:**
- Küçük harflere dönüştürür
- Büyük harflerin önüne alt çizgi ekler
- Boşlukları ve kısa çizgileri alt çizgi ile değiştirir
- Çift alt çizgileri kaldırır

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.SnakeCase("merhaba dünya")
	fmt.Println(result) // "merhaba_dünya"
	
	result = gosugar.SnakeCase("MerhabaDünya")
	fmt.Println(result) // "merhaba_dünya"
	
	result = gosugar.SnakeCase("merhaba-dünya")
	fmt.Println(result) // "merhaba_dünya"
}
```

---

### 20. `KebabCase(s string) string`

String'i kebab-case'e dönüştürür.

**Signature:**
```go
func KebabCase(s string) string
```

**Parametreler:**
- `s` (string): Dönüştürülecek string

**Dönüş Değeri:**
- `string`: kebab-case versiyonu

**Behavior:**
- snake_case'e benzer fakat alt çizgi yerine kısa çizgi kullanır

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.KebabCase("merhaba dünya")
	fmt.Println(result) // "merhaba-dünya"
	
	result = gosugar.KebabCase("MerhabaDünya")
	fmt.Println(result) // "merhaba-dünya"
	
	result = gosugar.KebabCase("merhaba_dünya")
	fmt.Println(result) // "merhaba-dünya"
}
```

---

### 21. `PascalCase(s string) string`

String'i PascalCase'e dönüştürür.

**Signature:**
```go
func PascalCase(s string) string
```

**Parametreler:**
- `s` (string): Dönüştürülecek string

**Dönüş Değeri:**
- `string`: PascalCase versiyonu

**Behavior:**
- camelCase'e benzer fakat ilk harfi de büyük yapar
- Her sözcük büyük harfle başlar

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.PascalCase("merhaba dünya")
	fmt.Println(result) // "MerhabaDünya"
	
	result = gosugar.PascalCase("merhaba-dünya")
	fmt.Println(result) // "MerhabaDünya"
	
	result = gosugar.PascalCase("merhaba_dünya")
	fmt.Println(result) // "MerhabaDünya"
}
```

---

### 22. `Slugify(s string) string`

String'i URL-uyumlu slug'a dönüştürür.

**Signature:**
```go
func Slugify(s string) string
```

**Parametreler:**
- `s` (string): Dönüştürülecek string

**Dönüş Değeri:**
- `string`: URL-uyumlu slug

**Behavior:**
- Küçük harflere dönüştürür
- Özel karakterleri kaldırır
- Boşlukları kısa çizgi ile değiştirir
- Sonundaki kısa çizgileri kaldırır
- URL'ler ve dosya adları için idealdir

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	result := gosugar.Slugify("Merhaba Dünya!")
	fmt.Println(result) // "merhaba-dünya"
	
	result = gosugar.Slugify("GoSugar 1.0 Yayınlandı!")
	fmt.Println(result) // "gosugar-10-yayınlandı"
	
	result = gosugar.Slugify("  İletişim Bilgileri  ")
	fmt.Println(result) // "iletisim-bilgileri"
}
```

---

## Örnekler

### Örnek 1: String Normalleştirme

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Form'dan kullanıcı input'u
	userInput := "  JoHn DoE  "
	
	// Normalize et
	name := gosugar.Trim(userInput)
	name = gosugar.Capitalize(name)
	
	fmt.Printf("Hoş geldiniz, %s!\n", name)
}
```

### Örnek 2: Değişken Adı Oluştur

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	titles := []string{
		"Kullanıcı Profili",
		"Ürün Detayları",
		"API Anahtarı",
	}
	
	for _, title := range titles {
		camel := gosugar.CamelCase(title)
		snake := gosugar.SnakeCase(title)
		pascal := gosugar.PascalCase(title)
		
		fmt.Printf("Başlık: %s\n", title)
		fmt.Printf("  camelCase: %s\n", camel)
		fmt.Printf("  snake_case: %s\n", snake)
		fmt.Printf("  PascalCase: %s\n", pascal)
		fmt.Println()
	}
}
```

### Örnek 3: URL Slug Oluşturma

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	articleTitles := []string{
		"GoSugar ile Başlarken",
		"CLI Uygulamaları Oluşturmak",
		"Go için En İyi Uygulamalar",
	}
	
	for _, title := range articleTitles {
		slug := gosugar.Slugify(title)
		url := fmt.Sprintf("https://blog.example.com/articles/%s", slug)
		fmt.Println(url)
	}
}
```

### Örnek 4: Metin Biçimlendirme

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Metni padding ile ortala
	text := "GoSugar"
	width := 20
	
	result := gosugar.PadLeft(text, (width+len(text))/2, ' ')
	result = gosugar.PadRight(result, width, ' ')
	
	fmt.Printf("|%s|\n", result)
	
	// Ayırıcı oluştur
	separator := gosugar.Repeat("=", 40)
	fmt.Println(separator)
}
```

### Örnek 5: Tablo Sütun Biçimlendirmesi

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Tablo başlığı
	headers := []string{"ID", "Ad", "Durum"}
	data := [][]string{
		{"1", "Kullanıcı1", "Aktif"},
		{"12", "UzunKullanıcıAdı", "İnaktif"},
		{"123", "Test", "Aktif"},
	}
	
	// Sütunları formatla
	for _, row := range data {
		for i, cell := range row {
			// Sütun genişliğine pad et
			padded := gosugar.PadRight(cell, 15, ' ')
			fmt.Print(padded)
		}
		fmt.Println()
	}
}
```

---

## İlişkili Modüller

- **`errors.go`**: Error handling
- **`input.go`**: Doğrulama ile kullanıcı input'u
- **`validators.go`**: Input doğrulama yardımcıları
