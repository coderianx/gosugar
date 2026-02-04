# API Referansı: env - Ortam Değişkenleri Yönetimi

Ortam değişkenleri uygulamanızın konfigürasyonunu kontrol eder (port, veritabanı URL'i, API anahtarları vb.). `env.go` modülü ortam değişkenlerini okumayı ve `.env` dosyalarını yüklemeyi kolaylaştırır.

## 📋 İçindekiler

- [Genel Bakış](#genel-bakış)
- [Fonksiyonlar](#fonksiyonlar)
- [Örnekler](#örnekler)
- [Tasarım Kararları](#tasarım-kararları)

---

## Genel Bakış

### Amaç

- `.env` dosyasından ortam değişkenlerini yüklemek
- String, Integer, Boolean değişkenleri belirtilen tip ile okumak
- Varsayılan (default) değerler sağlamak
- Zorunlu değişkenleri kontrol etmek

### Başlıca Özellikler

- ✅ `.env` dosya desteği (boş satırlar ve yorumlar göz ardı)
- ✅ Tür dönüşümü otomatik (string → int, bool)
- ✅ Varsayılan değer desteği
- ✅ Zorunlu değişken kontrol
- ✅ Varolan ortam değişkenlerini koruması (override etmez)
- ✅ Panic-based error handling (başarısız okumada panic)

### Workflow

```
.env dosyası okuma
    ↓
Anahtar=Değer çiftleri parse etme
    ↓
Ortam değişkenlerine setleme
    ↓
get fonksiyonları ile okuma (varsayılan ile)
```

---

## Fonksiyonlar

### 1. `EnvFile(path string)`

`.env` dosyasını okur ve tüm değişkenleri ortama yükler.

**Signature:**
```go
func EnvFile(path string)
```

**Parametreler:**
- `path` (string): `.env` dosyasının yolu

**Behavior:**
- Dosyayı açar ve satır satır okur
- `# ` ile başlayan satırları (yorum) göz ardı eder
- Boş satırları atlar
- `key=value` formatını parse eder
- **Varolan ortam değişkenlerini override etmez** (ortamda zaten varsa, skip eder)
- Hata varsa panic atar

**Hata Durumları:**
- Dosya bulunamadı: `panic("cannot open env file: ...")`
- Geçersiz format (`=` olmadan): `panic("invalid env line: ...")`
- Setenv başarısız: `panic("failed to set env ...")`

**Örnek `.env` dosyası:**
```env
# Uygulama Ayarları
APP_NAME=MyApp
PORT=8080
DEBUG=true

# Veritabanı
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=secret123
```

**Kullanım:**
```go
package main

import "github.com/coderianx/gosugar"

func main() {
	// .env dosyasını yükle
	gosugar.EnvFile(".env")
	
	// Artık değişkenleri okuyabilirsiniz
	appName := gosugar.EnvString("APP_NAME")  // "MyApp"
	port := gosugar.EnvInt("PORT")            // 8080
}
```

**İlişkili Fonksiyonlar:** `EnvString`, `EnvInt`, `EnvBool`, `MustEnv`

---

### 2. `EnvString(key string, defaultValue ...string) string`

String tipi ortam değişkeni okur.

**Signature:**
```go
func EnvString(key string, defaultValue ...string) string
```

**Parametreler:**
- `key` (string): Ortam değişkeninin adı
- `defaultValue` (variadic): Değişken yoksa dönecek varsayılan değer (opsiyonel)

**Dönüş Değeri:**
- Değişkenin değeri (yoksa: varsayılan değer veya boş string "")

**Behavior:**
- Ortamda `key` adında bir değişken var mı kontrol eder
- Varsa ve boş değilse, değerini döner
- Yoksa:
  - Varsayılan değer sağlandıysa, onu döner
  - Sağlanmadıysa, boş string "" döner
- **Panic atmaz** (hataya toleranslı)

**Hata Durumları:**
- Hiçbir hata hatası yok (her zaman başarılı döner)

**Örnekler:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Senaryo 1: Ortam değişkeni var
	os.Setenv("APP_NAME", "MyApp")
	name := gosugar.EnvString("APP_NAME", "DefaultApp")
	fmt.Println(name) // "MyApp"

	// Senaryo 2: Ortam değişkeni yok, varsayılan kullan
	theme := gosugar.EnvString("THEME", "dark")
	fmt.Println(theme) // "dark"

	// Senaryo 3: Ortam değişkeni yok, varsayılan da yok
	lang := gosugar.EnvString("LANG")
	fmt.Println(lang) // "" (boş string)

	// Senaryo 4: Ortam değişkeni boş string
	os.Setenv("EMPTY", "")
	val := gosugar.EnvString("EMPTY", "default")
	fmt.Println(val) // "default" (boş string varsayılan yerine döner)
}
```

**İlişkili Fonksiyonlar:** `EnvInt`, `EnvBool`, `MustEnv`

---

### 3. `EnvInt(key string, defaultValue ...int) int`

Integer tipi ortam değişkeni okur ve tür dönüşümü yapar.

**Signature:**
```go
func EnvInt(key string, defaultValue ...int) int
```

**Parametreler:**
- `key` (string): Ortam değişkeninin adı
- `defaultValue` (variadic): Dönüştürülemezse dönecek varsayılan değer

**Dönüş Değeri:**
- Dönüştürülmüş integer değer

**Behavior:**
- Ortam değişkenini okur
- `strconv.Atoi()` ile integer'a dönüştürür
- Dönüştürülemezse:
  - Varsayılan değer varsa, onu döner
  - Yoksa panic atar
- Değişken yoksa:
  - Varsayılan değer varsa, onu döner
  - Yoksa panic atar

**Hata Durumları:**
- Geçersiz format (integer değil): `panic("invalid int env var ...")`
- Değişken yok ve varsayılan yok: `panic("missing env var ...")`

**Örnekler:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Senaryo 1: Geçerli integer
	os.Setenv("PORT", "8080")
	port := gosugar.EnvInt("PORT")
	fmt.Println(port) // 8080

	// Senaryo 2: Yoksa varsayılan kullan
	timeout := gosugar.EnvInt("TIMEOUT", 30)
	fmt.Println(timeout) // 30

	// Senaryo 3: Geçersiz format, varsayılan düş
	os.Setenv("BAD_NUMBER", "not_a_number")
	value := gosugar.EnvInt("BAD_NUMBER", 0)
	fmt.Println(value) // 0

	// Senaryo 4: Geçersiz ve varsayılan yok → panic!
	// num := gosugar.EnvInt("MISSING_VAR") // panic!
}
```

**İlişkili Fonksiyonlar:** `EnvString`, `EnvBool`

---

### 4. `EnvBool(key string, defaultValue ...bool) bool`

Boolean tipi ortam değişkeni okur. Çeşitli string formatlarını destekler.

**Signature:**
```go
func EnvBool(key string, defaultValue ...bool) bool
```

**Parametreler:**
- `key` (string): Ortam değişkeninin adı
- `defaultValue` (variadic): Dönüştürülemezse dönecek varsayılan değer

**Dönüş Değeri:**
- Boolean değer (true veya false)

**Supported Values:**
- **True**: `"true"`, `"1"`, `"yes"`, `"y"`, `"on"` (büyük/küçük harf farketmez)
- **False**: `"false"`, `"0"`, `"no"`, `"n"`, `"off"` (büyük/küçük harf farketmez)

**Behavior:**
- Ortam değişkenini okur
- Değeri case-insensitive olarak kontrol eder
- Eğer tanınan bir format değilse:
  - Varsayılan değer varsa, onu döner
  - Yoksa panic atar

**Hata Durumları:**
- Geçersiz format: `panic("invalid bool env var ...")`
- Değişken yok ve varsayılan yok: `panic("missing env var ...")`

**Örnekler:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Senaryo 1: Standart true/false
	os.Setenv("DEBUG", "true")
	debug := gosugar.EnvBool("DEBUG")
	fmt.Println(debug) // true

	os.Setenv("PRODUCTION", "false")
	prod := gosugar.EnvBool("PRODUCTION")
	fmt.Println(prod) // false

	// Senaryo 2: Alternatif true değerleri
	os.Setenv("ENABLE_CACHE", "1")
	cache := gosugar.EnvBool("ENABLE_CACHE")
	fmt.Println(cache) // true

	os.Setenv("AUTO_START", "yes")
	auto := gosugar.EnvBool("AUTO_START")
	fmt.Println(auto) // true

	// Senaryo 3: Alternatif false değerleri
	os.Setenv("SKIP_VALIDATION", "0")
	skip := gosugar.EnvBool("SKIP_VALIDATION")
	fmt.Println(skip) // false

	// Senaryo 4: Varsayılan değer
	verbose := gosugar.EnvBool("VERBOSE", false)
	fmt.Println(verbose) // false

	// Senaryo 5: Geçersiz format, varsayılan
	os.Setenv("INVALID_BOOL", "maybe")
	value := gosugar.EnvBool("INVALID_BOOL", true)
	fmt.Println(value) // true
}
```

**İlişkili Fonksiyonlar:** `EnvString`, `EnvInt`

---

### 5. `MustEnv(key string) string`

**Zorunlu** ortam değişkeni okur. Değişken yoksa panic atar.

**Signature:**
```go
func MustEnv(key string) string
```

**Parametreler:**
- `key` (string): Ortam değişkeninin adı (zorunlu)

**Dönüş Değeri:**
- Ortam değişkeninin değeri (string)

**Behavior:**
- Ortamda `key` adında bir değişken var mı kontrol eder
- Varsa ve boş değilse, değerini döner
- Yoksa veya boşsa: **panic atar**
- Varsayılan değer desteklemez (kesinlikle gerekli)

**Hata Durumları:**
- Değişken yok: `panic("required env var missing: ...")`
- Değişken boş: `panic("required env var missing: ...")`

**Kullanım Senaryoları:**
- API anahtarları
- Veritabanı bağlantı stringi
- Kritik konfigürasyon değerleri

**Örnekler:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	// Senaryo 1: Değişken var
	os.Setenv("DATABASE_URL", "postgres://localhost/mydb")
	dbURL := gosugar.MustEnv("DATABASE_URL")
	fmt.Println(dbURL) // "postgres://localhost/mydb"

	// Senaryo 2: Değişken yok → panic!
	// apiKey := gosugar.MustEnv("API_KEY") // panic!

	// Senaryo 3: Değişken boş → panic!
	os.Setenv("EMPTY_VAR", "")
	// val := gosugar.MustEnv("EMPTY_VAR") // panic!
}
```

**Best Practice:**

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// Başlamada zorunlu değişkenleri kontrol et
	apiKey := gosugar.MustEnv("API_KEY")
	dbURL := gosugar.MustEnv("DATABASE_URL")
	
	// Opsiyonel değişkenler için varsayılan kullan
	port := gosugar.EnvInt("PORT", 8080)
	debug := gosugar.EnvBool("DEBUG", false)
	
	fmt.Printf("Başarıyla yüklendi: API_KEY, DATABASE_URL, port=%d\n", port)
}
```

**İlişkili Fonksiyonlar:** `EnvString`, `EnvInt`, `EnvBool`

---

## Örnekler

### Örnek 1: Basit Konfigürasyon

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	// .env dosyasını yükle
	gosugar.EnvFile(".env")

	// Konfigürasyonu oku
	appName := gosugar.EnvString("APP_NAME", "MyApp")
	port := gosugar.EnvInt("PORT", 8080)
	debug := gosugar.EnvBool("DEBUG", false)

	fmt.Printf("App: %s\n", appName)
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("Debug: %v\n", debug)
}
```

`.env`:
```env
APP_NAME=ProductionApp
PORT=3000
DEBUG=false
```

### Örnek 2: Zorunlu ve Opsiyonel Değişkenler

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
)

func main() {
	gosugar.EnvFile(".env")

	// Zorunlu değişkenler (yoksa panic)
	databaseURL := gosugar.MustEnv("DATABASE_URL")
	apiKey := gosugar.MustEnv("API_KEY")

	// Opsiyonel değişkenler (varsayılan ile)
	logLevel := gosugar.EnvString("LOG_LEVEL", "info")
	maxConnections := gosugar.EnvInt("MAX_CONNECTIONS", 10)
	enableCache := gosugar.EnvBool("ENABLE_CACHE", true)

	fmt.Printf("Database: %s\n", databaseURL[:20]+"...")
	fmt.Printf("Log Level: %s\n", logLevel)
	fmt.Printf("Max Conn: %d\n", maxConnections)
	fmt.Printf("Cache: %v\n", enableCache)
}
```

### Örnek 3: Environment'a Göre Davranış

```go
package main

import (
	"fmt"
	"github.com/coderianx/gosugar"
	"os"
)

func main() {
	gosugar.EnvFile(".env")

	env := gosugar.EnvString("ENVIRONMENT", "development")

	switch env {
	case "production":
		// Kesin zorunlu değişkenler
		_ = gosugar.MustEnv("DATABASE_URL")
		_ = gosugar.MustEnv("API_KEY")
		fmt.Println("Production mode: Tüm zorunlu değişkenler kontrol edildi")

	case "development":
		// Esnekçe kullanılabilir
		dbURL := gosugar.EnvString("DATABASE_URL", "localhost:5432")
		fmt.Printf("Dev mode: Database = %s\n", dbURL)

	default:
		fmt.Println("Bilinmeyen environment")
		os.Exit(1)
	}
}
```

---

## Tasarım Kararları

### 1. Neden Panic Seçildi?

`MustEnv` ve `EnvInt`/`EnvBool` (varsayılan olmadan) panic atar. Neden?

**Sebep:**
- Konfigürasyon hataları early-stage olmalı
- Uygulamanın başında kullanılır (startup validation)
- Hatalı konfigürasyon ile çalışmaktan daha iyidir

**Alternatif:** `Try` ile güvenli hale getirin:
```go
value, ok := gosugar.Try(func() string {
	return gosugar.MustEnv("CRITICAL_VAR")
})
```

### 2. Neden Varsayan Ortam Değişkenlerini Override Etmiyor?

Eğer ortamda zaten bir değişken varsa, `.env` dosyasından yüklenen değer kullanılmaz. Neden?

**Sebep:**
- Uygulamanın başında ortam değişkenleri setlenebilir
- Docker/Kubernetes containerlarında ENV'ler container başlatırken set edilir
- .env dosyası sadece "fallback" için kullanılır

**Sonuç:**
```bash
# Komut satırından başlat
PORT=9000 go run main.go

# Uygulamada EnvFile yapıldığında bile PORT=8080 değerini yok saymaz, 9000 kullanır
port := gosugar.EnvInt("PORT", 8080) // 9000
```

### 3. Neden Tür Dönüşümü Otomatik?

`EnvString` yerine `EnvInt` ve `EnvBool` ayrı fonksiyonlar neden var?

**Sebep:**
- Type safety: compile-time kontrol
- Hata yönetimi: dönüştürülemeyen değerler catch edilir
- Convenience: `.env` dosyasında `PORT=8080` yazıp direkt integer kullanmak

### 4. Boolean İçin Çok Değer Neden?

Boolean için `"true"`, `"1"`, `"yes"`, `"y"`, `"on"` neden?

**Sebep:**
- Farklı kültürler ve araçlar farklı formatlar kullanır
- Docker ve `docker-compose` `"1"`/`"0"` tercih eder
- Insan okunabilir: `"yes"`/`"no"` daha doğal

**Sonuç:**
```env
DEBUG=true        # Go style
CACHE_ENABLED=1   # Docker style
VERBOSE=yes       # Human readable
```

---

## Sık Sorulan Sorular

### P: `.env` dosyası production'da kullanılabilir mi?
**C:** Normalde hayır. Production'da ortam değişkenleri sistem ortamından ayarlanır (Docker ENV, Kubernetes secrets, system environment). `.env` sadece local development için.

### P: Birden fazla `.env` dosyası yükleyebilir miyim?
**C:** Evet, `EnvFile()` birden çok çağrılabilir:
```go
gosugar.EnvFile(".env")
gosugar.EnvFile(".env.local")  // İkinci dosya yüklenir
```
Ancak, ilk tanımlanan değerler korunur (override etmez).

### P: Boş satırlar ve yorumlar nasıl işlenir?
**C:** Boş satırlar (``) ve `#` ile başlayan satırlar otomatik atlanır:
```env
# Bu bir yorum

PORT=8080     # Bu da yorum

# DISABLED=true (bu atlanır)
```

### P: `EnvFile` dosya bulamazsa?
**C:** Panic atar. Hata toleranslı yapmak için:
```go
_, ok := gosugar.Try(func() {
	gosugar.EnvFile(".env.local")
})
```

### P: Specielçe eklenmiş ortam değişkenlerini görmek istiyorum
**C:** `os.Environ()` kullanın:
```go
for _, env := range os.Environ() {
	fmt.Println(env)
}
```

---

## Bağlantılı Modüller

- **`input.go`**: Kullanıcı inputu (env değişkenlerinden default değer setlenebilir)
- **`errors.go`**: Panic ve error handling
- **`file.go`**: Dosya okuma (EnvFile içinde kullanılır)

---

## Kaynaklar

- `env.go` kaynak kodu
- [`getting-started.md`](../guides/getting-started.md) - Başlama rehberi
- [`design-decisions.md`](../architecture/design-decisions.md) - Tasarım kararları

