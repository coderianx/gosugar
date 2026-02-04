# GoSugar - Başlangıç Rehberi

Bu rehber GoSugar ile ilk adımlarınızı atmanız için hazırlanmıştır. Kurulumdan ilk programınızı yazmanıza kadar her şeyi öğreneceksiniz.

## 🎯 Bu Rehberin Hedefi

Bu sayfayı bitirdikten sonra:
- ✅ GoSugar'ı yükleyebileceksiniz
- ✅ Temel fonksiyonları kullanabileceksiniz
- ✅ İlk CLI uygulamanızı yazabileceksiniz
- ✅ Nereye bakacağınızı bileceksiniz (sonraki adımlar)

**Okuma süresi:** ~15 dakika

---

## 1️⃣ Kurulum

### Ön Koşullar
- Go 1.18 veya daha yüksek
- Terminal/komut satırı (bash, zsh, cmd, PowerShell vb.)

### Kurulum Adımı

```bash
go get github.com/coderianx/gosugar
```

Bu komut GoSugar kütüphanesini indirir ve Go modülünüze ekler.

### Doğrulama

Kurulum başarılı mı kontrol etmek için basit bir test dosyası oluşturun:

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Test: Rastgele sayı
	num := gosugar.RandInt(1, 10)
	fmt.Println("Rastgele:", num)
}
```

Çalıştırın:
```bash
go run main.go
```

Çıktıda bir sayı görürseniz, kurulum tamam! ✅

---

## 2️⃣ Temel Konseptler

GoSugar 6 temel modülden oluşur:

### 📋 1. Ortam Değişkenleri (`env`)
Uygulamanızın konfigürasyonu (port, veritabanı URL'i vb.)

```go
gosugar.EnvString("APP_NAME", "MyApp")  // OK: "MyApp" veya ortam değeri
gosugar.EnvInt("PORT", 8080)            // OK: 8080 veya ortam değeri
gosugar.MustEnv("API_KEY")              // Zorunlu: yoksa panic atar
```

### ⌨️ 2. Kullanıcı Inputu (`input`)
Terminalde kullanıcıdan veri almak

```go
name := gosugar.Input("Adınız: ")           // String
age := gosugar.InputInt("Yaşınız: ", 18)   // Integer (default: 18)
price := gosugar.InputFloat("Fiyat: ", 0)  // Float (default: 0)
```

### ✔️ 3. Validatörler (`validators`)
Girdileri kontrol etmek

```go
email := gosugar.Input(
	"E-mail: ",
	gosugar.NotEmpty(),    // Boş olamaz
	gosugar.MinLen(5),     // En az 5 karakter
	gosugar.MaxLen(100),   // En fazla 100 karakter
)
```

### 🎲 4. Rastgele Veri (`random`)
Test ve demo için rastgele değerler

```go
dice := gosugar.RandInt(1, 6)              // 1-6 arası
random := gosugar.RandString(10)           // 10 karakter
options := []string{"A", "B", "C"}
choice := gosugar.Choice(options)          // Listeden seçim
```

### 🛡️ 5. Hata Yönetimi (`errors`)
Güvenli error handling

```go
file := gosugar.Must(os.Open("config.json"))       // Error varsa panic
gosugar.Check(someFunction())                       // Sadece error ıdırın
value, ok := gosugar.Try(riskyFunction)           // Güvenli çalıştırma
result := gosugar.Or(value, ok, defaultValue)    // Fallback ile
```

### 📁 6. Dosya İşlemleri (`file`)
Dosya okuma/yazma

```go
content := gosugar.ReadFile("data.txt")         // Oku
gosugar.WriteFile("output.txt", "Hello")       // Yaz
gosugar.CreateFile("new.txt", "Başlangıç")    // Oluştur (varsa skip)
gosugar.AppendFile("log.txt", "Log satırı\n")  // Ekle
```

---

## 3️⃣ İlk Uygulamanız

Şimdi küçük ama faydalı bir uygulama yazalım: **Basit Anket Uygulaması**

### Adım 1: Dosya Oluşturun

`survey.go` adında yeni bir dosya oluşturun:

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	fmt.Println("=== Hoş Geldiniz ===\n")

	// Kullanıcının adını al
	name := gosugar.Input(
		"Adınız: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(2),
	)

	// Yaşını al
	age := gosugar.InputInt("Yaşınız: ", 0)

	// E-mailini al
	email := gosugar.Input(
		"E-mail: ",
		gosugar.NotEmpty(),
		gosugar.MinLen(5),
	)

	// Sonuçları göster
	fmt.Println("\n=== Girdiğiniz Bilgiler ===")
	fmt.Printf("Ad: %s\n", name)
	fmt.Printf("Yaş: %d\n", age)
	fmt.Printf("E-mail: %s\n", email)
	fmt.Println("\nTeşekkürler!")
}
```

### Adım 2: Çalıştırın

```bash
go run survey.go
```

### Adım 3: Deneyim

```
=== Hoş Geldiniz ===

Adınız: John
Yaşınız: 25
E-mail: john@example.com

=== Girdiğiniz Bilgiler ===
Ad: John
Yaş: 25
E-mail: john@example.com

Teşekkürler!
```

**Tebrikler!** Ilk GoSugar uygulamanızı yazdınız! 🎉

---

## 4️⃣ Daha Karmaşık Örnek: Ortam Dosyası

GoSugar'ın temel gücü ortam yönetimidir. İşte bunu görelim:

### Adım 1: `.env` Dosyası Oluşturun

`.env` adında bir dosya oluşturun:

```env
# Uygulama Ayarları
APP_NAME=MyCLIApp
DEBUG=true
PORT=3000

# Veritabanı (örnek)
DB_HOST=localhost
DB_PORT=5432
```

### Adım 2: Kodu Yazın

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// .env dosyasını yükle
	gosugar.EnvFile(".env")

	// Ortam değişkenlerini oku
	appName := gosugar.EnvString("APP_NAME", "DefaultApp")
	debug := gosugar.EnvBool("DEBUG", false)
	port := gosugar.EnvInt("PORT", 8080)

	fmt.Printf("Uygulama: %s\n", appName)
	fmt.Printf("Debug: %v\n", debug)
	fmt.Printf("Port: %d\n", port)

	// İsteğe bağlı değişkenler
	theme := gosugar.EnvString("THEME", "dark")
	fmt.Printf("Tema: %s (varsayılan)\n", theme)
}
```

### Adım 3: Çalıştırın

```bash
go run main.go
```

Çıktı:
```
Uygulama: MyCLIApp
Debug: true
Port: 3000
Tema: dark (varsayılan)
```

**Önemli:** Ortam değişkenlerini `.env` dosyasından yüklemek, production'da güvenlidir ve konfigürasyon yönetimini kolaylaştırır.

---

## 5️⃣ Rastgele Veri ile Örnek: Mini Oyun

Basit bir "Sayı Tahmin Oyunu" yazalım:

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	fmt.Println("🎮 Sayı Tahmin Oyunu")
	fmt.Println("1-100 arası bir sayıyı tahmin et!\n")

	// Rastgele bir sayı seç (1-100)
	secretNumber := gosugar.RandInt(1, 100)
	attempts := 0
	maxAttempts := 7

	for attempts < maxAttempts {
		attempts++

		// Kullanıcıdan tahmin al
		guess := gosugar.InputInt(
			fmt.Sprintf("Deneme %d/%d - Tahmininiz: ", attempts, maxAttempts),
			0,
		)

		if guess == secretNumber {
			fmt.Printf("\n🎉 Bildin! Sayı %d idi. %d denemede başarılı!\n", secretNumber, attempts)
			return
		} else if guess < secretNumber {
			fmt.Println("📈 Daha yüksek bir sayı dene")
		} else {
			fmt.Println("📉 Daha düşük bir sayı dene")
		}
	}

	fmt.Printf("\n😢 Oyun bitti! Sayı %d idi.\n", secretNumber)
}
```

Çalıştırın ve oynayın! 🎮

---

## 6️⃣ Hata Yönetimi Örneği

GoSugar hataları panic ile yönetir. Bunu kullanırken dikkat edin:

### Güvenli Dosya Okuma

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Yol 1: Güvenli okuma (Try/Or ile)
	content, ok := gosugar.Try(func() string {
		return gosugar.ReadFile("config.json")
	})

	if !ok {
		fmt.Println("Dosya okunamadı, varsayılan kullanılıyor")
		content = "Varsayılan konfigürasyon"
	}

	fmt.Println(content)

	// Yol 2: Basit okuma (dosya yoksa panic)
	// data := gosugar.ReadFile("important.json")
	// Bu kodda dosya yoksa panic atar!
}
```

---

## 7️⃣ Sonraki Adımlar

Başarıyla temel konseptleri öğrendiniz! Sonraki adımlar:

### 📖 **Derinlemesine Öğrenmek İçin**
- Özel modüller için: [`../api/`](../api/) klasöründe referansları okuyun
- Mimariyiyi anlamak için: [`../architecture/ARCHITECTURE.md`](../architecture/ARCHITECTURE.md)

### 🛠️ **Pratik Yapmak İçin**
1. Kendi CLI uygulamanızı yazın
2. `.env` dosyası ile konfigürasyon yapın
3. Validatörler ile input kontrol edin
4. Hata yönetimini deneyin

### 🚀 **Özel Senaryolar**
- Error handling best practices: [`error-handling.md`](error-handling.md)
- Tasarım desenleri: [`design-patterns.md`](design-patterns.md)
- Test yazma: [`testing-with-gosugar.md`](testing-with-gosugar.md)

---

## ❓ Sık Sorulan Sorular

### P: GoSugar'ın web uygulamaları için kullanabilir miyim?
**C:** Evet! Ortam yönetimi, dosya işlemleri, hata yönetimi web uygulamalarında da işe yarar. Ancak input almak CLI için tasarlandığı için web'de doğrudan kullanılamaz.

### P: Hangi Go versiyonu gerekli?
**C:** Go 1.18+. GoSugar generics özelliğini kullanır (1.18'de tanıtıldı).

### P: Harici bağımlılık var mı?
**C:** Hayır! Sadece Go standart kütüphanesini kullanır.

### P: Panic atması tehlikeli mi?
**C:** Basit uygulamalarda sorun değil. Kritik sistemlerde, hataları catch etmek için `Try/Or` kullanın.

### P: Validatörler kendi validatör yazabilir miyim?
**C:** Evet! Bir fonksiyon yazmanız yeterli: [`design-patterns.md`](design-patterns.md) bak.

---

## 🎓 Bilgilendirme

Bu rehber şunları kapsamaz:
- ❌ Go program dilinin temel öğrenişi (for loop, variable vb.)
- ❌ Kütüphanenin tüm API'si (bkz. [`../api/`](../api/))
- ❌ İleri seviye kullanım (bkz. [`../architecture/`](../architecture/))

---

**İleri okumaya hazır mısınız?** Modül referanslarından birini seçin:
- 📋 [`../api/env.md`](../api/env.md) - Ortam değişkenleri derinlemesine
- ⌨️ [`../api/input.md`](../api/input.md) - Input derinlemesine
- 🎲 [`../api/random.md`](../api/random.md) - Rastgele veri derinlemesine
- 📁 [`../api/file.md`](../api/file.md) - Dosya işlemleri derinlemesine

Veya başka bir rehber:
- 🛡️ [`error-handling.md`](error-handling.md) - Error handling stratejileri
- 🏗️ [`design-patterns.md`](design-patterns.md) - Tasarım desenleri

Yazlı sorularınız varsa: `github.com/coderianx/gosugar` üzerinde issue açabilirsiniz! 🤝
