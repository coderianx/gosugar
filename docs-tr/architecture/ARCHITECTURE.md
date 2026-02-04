# GoSugar Mimarisi - Derinlemesine Açıklama

Bu dokümantasyon GoSugar kütüphanesinin tam mimarisini, tasarım kararlarını ve iç yapısını açıklar.

## 📋 İçindekiler

- [Genel Mimari](#genel-mimari)
- [Modül Tasarımı](#modül-tasarımı)
- [Veri Akışı](#veri-akışı)
- [Tasarım Prensipleri](#tasarım-prensipleri)
- [Bağımlılık Grafiği](#bağımlılık-grafiği)
- [Yaygın Patterns](#yaygın-patterns)

---

## Genel Mimari

### Single Package Design

GoSugar **single-package** mimarisi kullanır:

```
github.com/coderianx/gosugar/
├── env.go              # Modül 1
├── input.go            # Modül 2
├── validators.go       # Modül 3
├── random.go           # Modül 4
├── errors.go           # Modül 5
├── file.go             # Modül 6
├── http.go             # Modül 7
└── go.mod
```

**Avantajlar:**
- ✅ Simple import: `import "github.com/coderianx/gosugar"`
- ✅ Flat namespace: `gosugar.Input()`, `gosugar.RandInt()` vb.
- ✅ Kolay keşfedilebilirlik
- ✅ Az dosya sayısı

**Dezavantajlar:**
- ❌ Package büyüdüğünde namespace pollution
- ❌ Modülleri bağımsız kullanamıyorsunuz
- ❌ İç implementasyon detayları açık

### Paket Yapısı

```
package gosugar

// Tüm public fonksiyonlar
func EnvString(...) string
func Input(...) string
func RandInt(...) int
// ... etc
```

---

## Modül Tasarımı

### 1. Modül: `env.go` (Ortam Değişkenleri)

**Sorumluluk:**
- `.env` dosyası yükleme
- Ortam değişkenleri okuma (typed)
- Varsayılan değer sağlama

**Bağımlılıklar:**
- Go stdlib: `os`, `bufio`, `fmt`, `strconv`, `strings`

**Fonksiyonlar:**
```
EnvFile(path)              → .env dosyasından yükle
EnvString(key, default)    → String oku
EnvInt(key, default)       → Int oku (tip dönüşümü)
EnvBool(key, default)      → Bool oku (tip dönüşümü)
MustEnv(key)               → Zorunlu oku
```

**Workflow:**

```
User calls EnvFile(".env")
    ↓
Open file
    ↓
Scan lines
    ↓
Parse "key=value"
    ↓
Skip comments (#) and empty lines
    ↓
os.Setenv() (only if not exists)
    ↓
Done

User calls EnvString("PORT", 8080)
    ↓
os.LookupEnv("PORT")
    ↓
If exists and not empty: return value
If not exists/empty: return default
```

**Tasarım Kararları:**
- **Neden varolan değerleri override etmiyor?** Container ortamlarında (Docker) ENV'ler startup'ta set edilir. `.env` sadece fallback.
- **Neden string döndürüyor EnvInt/Bool?** Strconv'a dönüş yapıyor, hata yönetimi yapıyor.
- **Neden MustEnv yoksa panic?** Config hataları early-stage olmalı.

---

### 2. Modül: `input.go` (Kullanıcı Inputu)

**Sorumluluk:**
- Terminal'den kullanıcı inputu alma
- Tür dönüşümü (string → int, float)
- Validatörler uygulamak

**Bağımlılıklar:**
- Go stdlib: `bufio`, `fmt`, `os`, `strconv`, `strings`
- Internal: `validators.go` (Validator type)

**Fonksiyonlar:**
```
Input(prompt, validators...)     → String input
InputInt(prompt, default)        → Int input
InputFloat(prompt, default)      → Float input
inputRaw(prompt) → internal      → Raw string oku
```

**Workflow:**

```
User calls Input("Name: ", NotEmpty(), MinLen(3))
    ↓
inputRaw() → bufio.Scanner ile oku
    ↓
strings.TrimSpace()
    ↓
Her validator'ü çalıştır
    ↓
Validasyon başarısız: panic(error)
Validasyon başarılı: döner string
```

**Tasarım Kararları:**
- **Neden panic atar validasyon hatası?** Input validation başarısız olursa kullanıcı tekrar giriş yapmalı.
- **Neden separat Input/InputInt/InputFloat?** Type safety. Compile-time kontrol.

---

### 3. Modül: `validators.go` (Doğrulama)

**Sorumluluk:**
- Validator türü tanımlamak
- Hazır validatörler sağlamak
- Composable pattern

**Bağımlılıklar:**
- Go stdlib: `errors`, `fmt`

**Tipler ve Fonksiyonlar:**
```
type Validator func(string) error    // Type tanımı

NotEmpty() Validator                 // Hazır validatör
MinLen(n) Validator                  // Hazır validatör
MaxLen(n) Validator                  // Hazır validatör
```

**Functional Programming Pattern:**

```go
// Validator is a function type
type Validator func(string) error

// Returned function closes over 'n'
func MinLen(n int) Validator {
    return func(s string) error {
        if len(s) < n {
            return fmt.Errorf("minimum length is %d", n)
        }
        return nil
    }
}

// Zincirleme kullanım
validators := []Validator{
    NotEmpty(),
    MinLen(5),
    MaxLen(100),
}
```

**Tasarım Kararları:**
- **Neden function type?** Composable ve extensible. Kullanıcılar custom validator yazabilir.
- **Neden closure?** Parametreler (n) validator'a "embedded" olur.

---

### 4. Modül: `random.go` (Rastgele Veri)

**Sorumluluk:**
- Rastgele sayılar üretmek
- Rastgele string oluşturmak
- Listeden seçim yapmak

**Bağımlılıklar:**
- Go stdlib: `math/rand`, `time`

**Fonksiyonlar:**
```
init() func                          → Seed başlatma (auto)
RandInt(min, max) int               → [min, max] aralığında
RandFloat(min, max) float64         → [min, max) aralığında
RandBool() bool                     → 50/50 şans
RandString(length) string           → Harfler sadece
Choice[T](items []T) T              → Listeden seçim (generic)
```

**Tasarım Kararları:**
- **Neden init()?** Seed otomatik başlatılır. Her run farklı random.
- **Neden RandInt inclusive, RandFloat exclusive?** Go stdlib pattern (math/rand.Intn exclusive, Float64 [0,1))
- **Neden Choice generic?** Type-safe. String, int, struct vb. her şeyle çalışır.

---

### 5. Modül: `errors.go` (Hata Yönetimi)

**Sorumluluk:**
- Panic patterns (Must, Check)
- Panic recovery (Try)
- Fallback mekanizması (Or)

**Bağımlılıklar:**
- Go stdlib: (none directly, sadece built-in defer/recover)

**Fonksiyonlar:**
```
Must[T](v T, err) T                 → err varsa panic
Check(err)                          → err varsa panic
Try[T](fn func() T) (T, bool)       → Panic recover
Or[T](v, ok, fallback) T            → Ternary-like
Ignore(err)                         → Error'u yut
```

**Workflow:**

```
Must Pattern:
    file, err := os.Open("file.txt")
    f := gosugar.Must(file, err)    // err varsa panic
    ↓
    if err != nil {
        panic(err)
    }
    return file

Try/Or Pattern:
    value, ok := gosugar.Try(func() int {
        return 100 / 0              // panic
    })
    ↓
    defer recover catches panic
    ↓
    ok = false, value = zero-value
    ↓
    result := gosugar.Or(value, ok, 0) // 0 döner
```

**Tasarım Kararları:**
- **Neden panic?** Başarısız olan kritik operasyonlar program'ı durdurmalı.
- **Neden Try/Or?** Non-kritik operasyonlar graceful fallback.
- **Nedir Ignore?** Linter warnings'ı kaldırmak: `_ = err` yerine.

---

### 6. Modül: `file.go` (Dosya İşlemleri)

**Sorumluluk:**
- Dosya okuma
- Dosya yazma
- Dosya oluşturma (protective)
- Dosya ekleme (append)

**Bağımlılıklar:**
- Go stdlib: `fmt`, `os`

**Fonksiyonlar:**
```
ReadFile(path) string               → Oku
WriteFile(path, content)            → Yaz (overwrite)
CreateFile(path, content)           → Oluştur (varsa skip)
AppendFile(path, content)           → Ekle (yoksa oluştur)
```

**Tasarım Kararları:**
- **Neden CreateFile varsa skip?** Template dosyaları, varsayılan configs korumak için.
- **AppendFile yoksa oluştur?** Logging'de çok yaygın. Her call'da kontrol yapılmasın.

---

### 7. Modül: `http.go` (HTTP İstekleri)

**Sorumluluk:**
- HTTP GET istekleri
- Response body okuma
- JSON decode
- Headers okuma

**Bağımlılıklar:**
- Go stdlib: `encoding/json`, `fmt`, `io`, `net/http`

**Fonksiyonlar:**
```
GetBody(url) (string, error)               → Body oku
MustGetBody(url) string                    → Body oku (panic)
GetJSON[T](url) (T, error)                 → JSON decode
GetHeader(url) (http.Header, error)        → Headers
MustGetHeader(url) http.Header             → Headers (panic)
```

---

## Veri Akışı

### Senaryo 1: CLI Uygulaması

```
┌──────────────────────────────┐
│    Uygulama Başlanır         │
└──────────────┬───────────────┘
               │
               ▼
    ┌─────────────────────┐
    │ env.go              │
    │ EnvFile(".env")     │
    │ EnvString(...)      │ ──→ ortam değişkenleri
    └─────────────────────┘
               │
               ▼
    ┌─────────────────────┐
    │ input.go            │
    │ Input("Q: ")        │ ──→ validatörler (validators.go)
    │ InputInt(...)       │
    └─────────────────────┘
               │
               ▼
    ┌──────────────────────────┐
    │ İşleme Logik            │
    │ (Kullanıcı kodu)        │
    │ - random data gen       │ ──→ random.go
    │ - file I/O              │ ──→ file.go
    │ - error handling        │ ──→ errors.go
    └──────────────────────────┘
               │
               ▼
    ┌─────────────────────┐
    │ file.go             │
    │ WriteFile(...)      │ ──→ Sonuç dosyası
    │ AppendFile(...)     │
    └─────────────────────┘
```

### Senaryo 2: API İletişimi

```
┌──────────────────────────┐
│ API Kodu                 │
└──────────┬───────────────┘
           │
           ▼
    ┌──────────────────────┐
    │ http.go              │
    │ GetJSON[T](url)      │
    └──────────┬───────────┘
               │
               ▼
    ┌──────────────────────┐
    │ errors.go (implicit) │
    │ HTTP error ise       │
    │ error döner          │
    └──────────┬───────────┘
               │
               ▼
    ┌──────────────────────┐
    │ Kullanıcı Kodu       │
    │ Try/Or pattern ile   │
    │ fallback sağla       │ ──→ errors.go
    └──────────────────────┘
```

---

## Tasarım Prensipleri

### 1. **Simplicity First**

Go stdlib'ı wrapping, replacing değil:

```go
// ✅ GoSugar - wrapper
func EnvString(key, default) string {
    return os.LookupEnv(key) // stdlib'ı wrap
}

// ❌ Replacement (çok karmaşık)
// custom environment variable system
```

### 2. **Zero Dependencies**

Sadece Go stdlib:

```go
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    // ... sadece stdlib
)

// ❌ NO external packages
// import "github.com/some/package"
```

### 3. **Panic-Based Error Handling**

Simple uygulamalar için:

```go
// ✅ Simple apps
apiKey := gosugar.MustEnv("API_KEY") // yoksa panic

// ❌ Production apps (dikkat gerekli)
value, ok := gosugar.Try(someRiskyOp) // recover ederim
```

### 4. **Type Safety (Generics)**

Go 1.18+ generics:

```go
// ✅ Type-safe
choice := gosugar.Choice([]string{"A", "B"}) // string
num := gosugar.Choice([]int{1, 2})            // int

// ❌ Type-unsafe (any conversion)
choice := someChoice([]interface{}{...})
```

### 5. **Composability**

Validatörler zincirlenebilir:

```go
// ✅ Zincirli validatörler
Input(
    "E-mail: ",
    NotEmpty(),
    MinLen(5),
    MaxLen(100),
)

// ❌ Tüm validasyonu elle
if email == "" { ... }
if len(email) < 5 { ... }
```

---

## Bağımlılık Grafiği

```
input.go ──────┐
               ├──→ validators.go
               │
               ├──→ errors.go (implicit panics)
               │
               └──→ Go stdlib

env.go ────────┐
               ├──→ errors.go (panic)
               │
               └──→ Go stdlib

file.go ───────┐
               ├──→ errors.go (panic)
               │
               └──→ Go stdlib

random.go ─────→ Go stdlib (bağımsız)

http.go ───────┐
               ├──→ errors.go (implicit via error handling)
               │
               └──→ Go stdlib

errors.go ─────→ Go stdlib (bağımsız)

validators.go ─→ Go stdlib (bağımsız)
```

**Derinlik:** Maksimum 2 level (circular dependency yok)

---

## Yaygın Patterns

### Pattern 1: Startup Configuration

```go
func main() {
    // Load config
    gosugar.EnvFile(".env")
    
    // Validate critical vars
    apiKey := gosugar.MustEnv("API_KEY")      // yoksa panic
    port := gosugar.EnvInt("PORT", 8080)      // default ile
    
    // Start app
    startServer(port, apiKey)
}
```

### Pattern 2: User Input Loop

```go
for {
    // Validatörlerle input al
    command := gosugar.Input(
        "> ",
        gosugar.NotEmpty(),
    )
    
    // İşle
    handleCommand(command)
    
    // Validasyon başarısız olursa tekrar sor
}
```

### Pattern 3: Graceful Fallback

```go
// Ana kaynak dene
config, err := getConfigFromAPI()
if err == nil {
    // Başarılı
    useConfig(config)
} else {
    // Fallback: local dosya
    config = gosugar.ReadFile("config.local.json")
    useConfig(config)
}

// Ya da Try/Or
config, ok := gosugar.Try(getConfigFromAPI)
config = gosugar.Or(config, ok, defaultConfig)
```

### Pattern 4: Test Data Generation

```go
func generateTestData(count int) {
    for i := 0; i < count; i++ {
        user := User{
            ID:   gosugar.RandInt(1, 10000),
            Name: gosugar.RandString(10),
            Active: gosugar.RandBool(),
        }
        saveUser(user)
    }
}
```

---

## Genişletilebilirlik

GoSugar tasarımı genişlemeyi teşvik eder:

### Kendi Validatörü Yazma

```go
func IsEmail() gosugar.Validator {
    return func(s string) error {
        if !strings.Contains(s, "@") {
            return fmt.Errorf("not an email")
        }
        return nil
    }
}

// Kullan
email := gosugar.Input("E-mail: ", IsEmail())
```

### Kendi Error Pattern'ı

```go
// Try/Or yerine custom pattern
if err := riskyOperation(); err != nil {
    log.Printf("Operation failed: %v", err)
    handleError(err)
}
```

### Yeni Modül Ekleme

Yeni modülü `packagename/modulename.go` olarak ekle:

```go
// Örnek: string.go
package gosugar

func Reverse(s string) string {
    // implementation
}
```

---

## Sonuç

GoSugar mimarisi:
- 📦 **Single-package** design (simple)
- 🎯 **Focused** fonksiyonlar (her modül bir şeye odaklanır)
- 🔗 **Minimal coupling** (modüller bağımsız)
- 🛡️ **Panic-based** error handling (simple apps için)
- 💪 **Extensible** (custom validatör, pattern vb.)

Next: [`design-decisions.md`](design-decisions.md) tasarım kararlarının detayları için.

