# Tasarım Kararları - Neden Böyle Tasarlandı?

Bu dokümantasyon GoSugar'ın **neden** bu şekilde tasarlandığını açıklar. Her karar, trade-off'ları ve alternatiflerle birlikte.

## 📋 İçindekiler

- [Neden Panic?](#neden-panic)
- [Neden Single Package?](#neden-single-package)
- [Neden Generics?](#neden-generics)
- [Neden Fonksiyon Tipleri?](#neden-fonksiyon-tipleri)
- [Neden Zero Dependencies?](#neden-zero-dependencies)
- [Sık Sorular](#sık-sorular)

---

## Neden Panic?

### Karar: Hata yönetimi için panic kullanılıyor

**Kod Örneği:**
```go
// GoSugar
func MustEnv(key string) string {
    value, ok := os.LookupEnv(key)
    if !ok || value == "" {
        panic(fmt.Errorf("required env var missing: %s", key))  // ← PANIC
    }
    return value
}
```

### Neden?

| Durum | Panic İle | Error İle |
|-------|-----------|-----------|
| **Konfigürasyon hatası** | ✅ Hemen öğren | ❌ Çalışmaya devam et (yanlış) |
| **Startup validation** | ✅ Temiz | ❌ Her yerde kontrol gerek |
| **Code clarity** | ✅ Basit | ❌ Karışık error handling |
| **Simple apps** | ✅ Uygun | ❌ Gereksiz yapı |

### Karar Verme

```
Hedef Kitle: CLI uygulamaları, scripts, small projects
           ↓
Yanıt: Panic uygun
           ↓
Gerekçe: Config hataları erken catch edilmeli
```

### Alternatifleri

**Alternatif 1: Error Döndürme**

```go
// ❌ Daha verbose
func EnvString(key string) (string, error) {
    value, ok := os.LookupEnv(key)
    if !ok {
        return "", fmt.Errorf("missing: %s", key)
    }
    return value, nil
}

// Kullanımı
value, err := gosugar.EnvString("KEY")
if err != nil {
    // ... error handling
}
```

**Çıkarım:** Production-grade error handling için `Try/Or` patterns'ı varır.

---

## Neden Single Package?

### Karar: Tüm fonksiyonlar `gosugar` paketinde

**Yapı:**
```
gosugar/
├── env.go
├── input.go
├── validators.go
├── random.go
├── errors.go
├── file.go
└── http.go

# NOT: subdirectories yok
```

**Kullanım:**
```go
import "github.com/coderianx/gosugar"

gosugar.EnvString("KEY")
gosugar.Input("Q: ")
gosugar.RandInt(1, 10)
// All in one place
```

### Neden?

| Aspekt | Single Package | Multiple Packages |
|--------|---|---|
| **Import** | `import "...gosugar"` | `import "...gosugar/env"` `import "...gosugar/input"` |
| **Namespace** | `gosugar.Func()` | `env.Func()`, `input.Func()` |
| **Discovery** | ✅ Hepsi bir yerde | ❌ Hangisi nerede? |
| **Dependencies** | 🟡 Tüm modülleri yükle | ✅ Lazım olanı yükle |
| **Simple apps** | ✅ Minimal import | ❌ Birden fazla import |

### Alternatifleri

**Alternatif 1: Subpackages**

```
gosugar/
├── env/
│   └── env.go
├── input/
│   └── input.go
└── go.mod
```

**Kullanım:**
```go
import (
    "github.com/coderianx/gosugar/env"
    "github.com/coderianx/gosugar/input"
)

env.EnvString(...)
input.Input(...)
```

**Pros:**
- ✅ Isteğe bağlı import
- ✅ Namespace organization

**Cons:**
- ❌ Daha karmaşık
- ❌ Validator'ler nereye? (shared?)

**Alternatif 2: Monolithic File**

```
gosugar.go (1000+ lines)
```

**Cons:**
- ❌ Okuması zor
- ❌ Bağımlılıkları takip zor

---

## Neden Generics?

### Karar: Go 1.18+ generics kullanılıyor

**Kod Örneği:**
```go
// ✅ GoSugar - Generic
func Choice[T any](items []T) T {
    return items[rand.Intn(len(items))]
}

// Kullanım
fruit := gosugar.Choice([]string{"apple", "banana"})
num := gosugar.Choice([]int{1, 2, 3})
```

### Neden?

| Avantaj | Açıklama |
|---------|----------|
| **Type Safety** | Compile-time kontrolü. İnsan hatası azalır. |
| **No Casting** | `interface{}` dönüştürme gerekli değil |
| **Performance** | Runtime reflection yok. Hızlı. |
| **Clarity** | Intent açık. Kodu okuması kolay. |

**Örnek: Generics Olmadan**

```go
// ❌ Go 1.17 - interface{}
func Choice(items []interface{}) interface{} {
    return items[rand.Intn(len(items))]
}

// Kullanımı
data := gosugar.Choice([]interface{}{"a", "b"})
fruit := data.(string)  // ← Type assertion gerekli
```

### Alternatifleri

**Alternatif 1: Type-Specific Fonksiyonlar**

```go
// ❌ Tekrar
func ChoiceString(items []string) string { ... }
func ChoiceInt(items []int) int { ... }
func ChoiceFloat(items []float64) float64 { ... }
```

**Cons:**
- ❌ Çok fazla fonksiyon
- ❌ DRY prensibine aykırı

---

## Neden Fonksiyon Tipleri?

### Karar: Validatörler fonksiyon tipi olarak

**Kod Örneği:**
```go
// ✅ GoSugar - Function Type
type Validator func(string) error

func MinLen(n int) Validator {
    return func(s string) error {
        if len(s) < n {
            return fmt.Errorf("min %d", n)
        }
        return nil
    }
}
```

### Nedir Bu?

**Functional Programming Pattern:**

```
MinLen(5) çağrısı:
    ↓
Bir fonksiyon döner (closure with captured 'n')
    ↓
Input("Q: ", MinLen(5)) çağrısında kullanılır
    ↓
Her input için validator fonksiyonu çalışır
```

### Neden?

| Sebep | Açıklama |
|-------|----------|
| **Composability** | Validatörler zincirlenebilir |
| **Flexibility** | Kullanıcılar custom validator yazabilir |
| **Simplicity** | Interface'ten daha basit |
| **Higher-Order Funcs** | FP pattern'ı modern Go'da |

**Örnek: Zincir**

```go
// Input her validatörü çalıştırır
Input(
    "Q: ",
    NotEmpty(),       // validator 1
    MinLen(5),        // validator 2
    MaxLen(100),      // validator 3
)
```

### Alternatifleri

**Alternatif 1: Interface**

```go
// ❌ Aşırı mühendislik
type Validator interface {
    Validate(string) error
}

type NotEmptyValidator struct{}
func (n NotEmptyValidator) Validate(s string) error { ... }

type MinLenValidator struct{ n int }
func (m MinLenValidator) Validate(s string) error { ... }

// Kullanım
Input("Q: ", NotEmptyValidator{}, MinLenValidator{5})
```

**Cons:**
- ❌ Fazla kod
- ❌ Boilerplate

**Alternatif 2: Struct avec Methods**

```go
// ❌ Daha fazla durum
type InputValidator struct {
    NotEmpty bool
    MinLen   int
    MaxLen   int
}
```

---

## Neden Zero Dependencies?

### Karar: Sadece Go stdlib

**go.mod:**
```go
module github.com/coderianx/gosugar

go 1.25.5

// No require statements!
```

### Neden?

| Sebep | Etkileri |
|-------|----------|
| **Simplicity** | Başlamak kolay. Karmaşık setup yok. |
| **Stability** | External package güncellemeleri sorun yaratmaz |
| **Size** | Binary küçük |
| **Production** | Deploy'da minimal risk |
| **Learning** | Go stdlib'ı öğren. Best practice'ler al. |

### Trade-offs

| Senaryo | Zero Deps | With Deps |
|--------|-----------|-----------|
| **String manipulation** | ✅ stdlib yeterli | ❌ `github.com/urfave/cli` |
| **HTTP requests** | ✅ net/http | ❌ `github.com/go-resty/resty` |
| **JSON** | ✅ encoding/json | ❌ `github.com/json-iterator/go` |

---

## Sık Sorular

### P: GoSugar'ı production'da kullanabilir miyim?

**C:** Kısmen:
- ✅ **Genel utility'ler:** `Input`, `RandInt`, `File` OK
- ✅ **Config management:** `env` OK
- ⚠️ **Error handling:** Panic çok agresif, `Try/Or` kullan
- ❌ **High-frequency ops:** HTTP, Database yok

**Tavsiye:**
```go
// Production: sadece config yönetimi
gosugar.EnvFile(".env")
port := gosugar.EnvInt("PORT", 8080)

// Non-critical: CLI input
name := gosugar.Input("Name: ")

// Kritik: stdlib kullan
db, err := sql.Open(...)
if err != nil { /* proper error handling */ }
```

### P: Neden panic'ten önce Try/Or değil?

**C:** UX açısından:
- Panic atarsa: **hata mesajı açık**
- Try/Or: **hata sessizce geçilir** (fark etmek zor)

```go
// Panic: açık
port := gosugar.EnvInt("PORT")  // CRASH, clear message

// Try/Or: sakın
value, ok := gosugar.Try(func() int {
    return gosugar.EnvInt("PORT")
})
port := gosugar.Or(value, ok, 8080)  // Sessizce 8080 kullan
```

### P: Neden HTTP modülü sınırlı?

**C:** İlaçlı HTTP başlı başına kompleks:
- Custom headers
- POST/PUT/DELETE
- Timeout
- Retry logic
- Authentication

**Karar:** Basit GET çağrılarını support et, REST client'lardan kaçın.

**Çıkarım:** `net/http` package'ını direkt kullanın (daha iyi).

### P: Özel use-case'im var. Genişletebilir miyim?

**C:** **Evet!** Kendi validatörü, kendi pattern'ı yaz:

```go
// Custom validator
func NumericOnly() gosugar.Validator {
    return func(s string) error {
        for _, ch := range s {
            if ch < '0' || ch > '9' {
                return fmt.Errorf("only numbers")
            }
        }
        return nil
    }
}

// Kullan
code := gosugar.Input("Code: ", NumericOnly())
```

### P: Neden dil Turkish?

**C:** Info.md Türkçeydi. Consistency için Türkçe devam etmek istenmiş.

---

## Sonuç

GoSugar tasarım felsefesi:

```
┌─────────────────────────────────────┐
│  Simplicity > Flexibility            │
│  Clarity > Performance (95% cases)   │
│  Single Package > Modular Packages   │
│  Type Safety > Dynamic              │
│  Zero Deps > Feature Completeness   │
└─────────────────────────────────────┘
```

**Kime uygun?**
- ✅ CLI uygulamaları
- ✅ Scripts ve automation
- ✅ Prototyping
- ✅ Learning Go

**Kime uygun değil?**
- ❌ Enterprise systems
- ❌ High-performance apps
- ❌ Complex business logic

---

## İlgili Dosyalar

- [`ARCHITECTURE.md`](ARCHITECTURE.md) - Teknik mimari
- [`../guides/design-patterns.md`](../guides/design-patterns.md) - Kullanım patterns'ları

