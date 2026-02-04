# API Referansı: file - Dosya İşlemleri

Dosya okuma, yazma, oluşturma ve ekleme işlemlerini basitleştiren modül.

## 📋 İçindekiler

- [Genel Bakış](#genel-bakış)
- [Fonksiyonlar](#fonksiyonlar)
- [Örnekler](#örnekler)

---

## Genel Bakış

### Amaç

- Dosya okumayı basitleştirmek
- Dosya yazma işlemlerini kolaylaştırmak
- Dosya oluşturma ve ekleme operasyonları

### Başlıca Özellikler

- ✅ UTF-8 string desteği
- ✅ Otomatik error handling (panic)
- ✅ CreateFile varlı dosyaları korur
- ✅ AppendFile yoksa dosya oluşturur

---

## Fonksiyonlar

### 1. `ReadFile(path string) string`

Dosyayı okur ve içeriğini string olarak döner.

**Signature:**
```go
func ReadFile(path string) string
```

**Parametreler:**
- `path` (string): Dosya yolu

**Dönüş Değeri:**
- Dosya içeriği (string)

**Behavior:**
- `os.ReadFile()` kullanır
- Tüm içeriği memory'ye yükler
- Hata varsa panic atar

**Hata Durumları:**
- Dosya bulunamadı: `panic("cannot read file ...")`
- Permission hatası: `panic("cannot read file ...")`

**Örnek:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Dosya oku
	content := gosugar.ReadFile("data.txt")
	fmt.Println(content)

	// JSON oku
	config := gosugar.ReadFile("config.json")
	fmt.Println("Config:", config)
}
```

---

### 2. `WriteFile(path string, content string)`

Dosyaya yazma yapır. Varsa üstüne yazar, yoksa oluşturur.

**Signature:**
```go
func WriteFile(path string, content string)
```

**Parametreler:**
- `path` (string): Dosya yolu
- `content` (string): Yazılacak içerik

**Behavior:**
- `os.WriteFile()` kullanır
- 0644 permissions ile oluşturur
- Varsa içeriği tamamen değiştirir (append değil!)
- Hata varsa panic atar

**Hata Durumları:**
- Permission hatası: `panic("cannot write file ...")`
- Geçersiz path: `panic("cannot write file ...")`

**Örnek:**

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	// Yeni dosya yaz
	gosugar.WriteFile("output.txt", "Hello World!")

	// Varsa üstüne yaz
	gosugar.WriteFile("output.txt", "Updated content")
}
```

**Uyarı:** Varsa önceki içerik silinir!

---

### 3. `CreateFile(path string, content string)`

Dosya **yoksa** oluşturur. Varsa hiçbir şey yapmaz.

**Signature:**
```go
func CreateFile(path string, content string)
```

**Parametreler:**
- `path` (string): Dosya yolu
- `content` (string): İlk içerik

**Behavior:**
- Dosya yoksa: oluşturur ve içeriği yazır
- Dosya varsa: hiçbir şey yapmaz (sessizce çıkar)
- 0644 permissions ile oluşturur

**Hata Durumları:**
- CreateFile başarısız: `panic("cannot create file ...")`
- Diğer hatalar: `panic("cannot check file ...")`

**Örnek:**

```go
package main

import "github.com/cosugar"

func main() {
	// İlk kez: oluşturur
	gosugar.CreateFile("config.json", "{\"port\": 8080}")

	// İkinci kez: yapılmaz
	gosugar.CreateFile("config.json", "{\"port\": 3000}")

	// Sonuç: config.json hala "{\"port\": 8080}" içeriğine sahip
}
```

**Best Practice:** Template dosyalar, default konfigürasyonlar için ideal.

---

### 4. `AppendFile(path string, content string)`

Dosyaya **ekleme** yapar. Yoksa oluşturur.

**Signature:**
```go
func AppendFile(path string, content string)
```

**Parametreler:**
- `path` (string): Dosya yolu
- `content` (string): Eklenecek içerik

**Behavior:**
- Dosya varsa: sonuna ekleme yapır
- Dosya yoksa: oluşturur ve içeriği yazır
- 0644 permissions ile oluşturur
- Mevcut içeriği değiştirmez

**Hata Durumları:**
- Append başarısız: `panic("cannot append to file ...")`

**Örnek:**

```go
package main

import "github.com/coderianx/gosugar"

func main() {
	// Log yazma
	gosugar.AppendFile("app.log", "Server started\n")
	gosugar.AppendFile("app.log", "Connection established\n")
	gosugar.AppendFile("app.log", "User logged in\n")

	// Sonuç: app.log tüm satırları içerir
}
```

---

## Örnekler

### Örnek 1: Konfigürasyon Dosyası

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Default config oluştur (varsa yapılmaz)
	defaultConfig := `{
	"app_name": "MyApp",
	"port": 8080,
	"debug": false
}`
	gosugar.CreateFile("config.json", defaultConfig)

	// Config oku
	config := gosugar.ReadFile("config.json")
	fmt.Println("Konfigürasyon:")
	fmt.Println(config)
}
```

### Örnek 2: Log Sistemi

```go
package main

import (
	"fmt"
	"time"
	"github.com/coderianx/gosugar"
)

func main() {
	logFile := "app.log"

	// Log yazma fonksiyonu
	writeLog := func(level, message string) {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		entry := fmt.Sprintf("[%s] %s: %s\n", timestamp, level, message)
		gosugar.AppendFile(logFile, entry)
	}

	// Örnek log'lar
	writeLog("INFO", "Uygulama başladı")
	writeLog("DEBUG", "Database bağlantısı açılıyor")
	writeLog("INFO", "Database bağlantısı başarılı")
	writeLog("ERROR", "API anahtarı bulunamadı")

	// Log'ları oku
	logs := gosugar.ReadFile(logFile)
	fmt.Println("=== Günlük ===")
	fmt.Println(logs)
}
```

### Örnek 3: İçerik İşleme

```go
package main

import (
	"fmt"
	"strings"
	"github.com/coderianx/gosugar"
)

func main() {
	// Dosya oku
	content := gosugar.ReadFile("input.txt")

	// İşle
	lines := strings.Split(content, "\n")
	fmt.Printf("Satır sayısı: %d\n", len(lines))

	// Yaz
	result := strings.Join(lines, " ")
	gosugar.WriteFile("output.txt", result)

	fmt.Println("İşleme tamamlandı")
}
```

### Örnek 4: Veri Dışa Aktarma

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Export başla
	reportFile := "report.csv"

	// Header
	gosugar.CreateFile(reportFile, "ID,Name,Score\n")

	// Veri ekle
	for i := 1; i <= 5; i++ {
		entry := fmt.Sprintf("%d,User%d,%.2f\n", i, i, float64(i)*10.5)
		gosugar.AppendFile(reportFile, entry)
	}

	// Rapor oku ve göster
	report := gosugar.ReadFile(reportFile)
	fmt.Println("=== Rapor ===")
	fmt.Println(report)
}
```

---

## Best Practices

### 1. CreateFile için Template Dosyalar

```go
// Ilk kez çalışırken varsayılan dosya oluştur
defaultEnv := `APP_NAME=MyApp
PORT=8080
DEBUG=false`

gosugar.CreateFile(".env", defaultEnv)
```

### 2. AppendFile için Logging

```go
// Loglama yapması gereken her yerde
gosugar.AppendFile("debug.log", "Operation started\n")
// ... işlem yapılır ...
gosugar.AppendFile("debug.log", "Operation completed\n")
```

### 3. WriteFile ile Overwrite

```go
// İçeriği tamamen değiştirmek istiyorsanız
newContent := processData(oldContent)
gosugar.WriteFile("processed.txt", newContent)
```

### 4. Error Handling (Try/Or ile)

```go
// Eğer dosya yoksa başarısız olabilir
content, ok := gosugar.Try(func() string {
	return gosugar.ReadFile("optional.txt")
})

data := gosugar.Or(content, ok, "default content")
```

---

## İlişkili Modüller

- **`errors.go`**: Error handling (Try/Or)
- **`env.go`**: .env dosyası yükleme (EnvFile)
- **`getting-started.md`**: Başlama rehberi

