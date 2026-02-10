# API Referansı: logger - Günlükleme Aracı

`logger.go` modülü, uygulamanızın çıktısını Debug'dan Fatal'a kadar önem düzeyleriyle kontrol etmenize yardımcı olan basit, seviye tabanlı bir günlükleme sistemi sağlar.

## 📋 İçerik

- [Genel Bakış](#genel-bakış)
- [Fonksiyonlar](#fonksiyonlar)
- [Günlük Seviyeleri](#günlük-seviyeleri)
- [Örnekler](#örnekler)
- [Tasarım Kararları](#tasarım-kararları)

---

## Genel Bakış

### Amaç

- Zaman damgası ile yapılandırılmış günlük mesajları oluşturmak
- Günlük seviyelerini kullanarak günlükleme ayrıntılılığını kontrol etmek
- Farklı önem düzeyleri için farklı türde mesajları desteklemek
- Ölümcül hatalar üzerinde uygulamayı kapatmak

### Ana Özellikler

- ✅ Beş günlük seviyesi (Debug, Info, Warn, Error, Fatal)
- ✅ Her günlük mesajı için zaman damgası biçimlendirmesi
- ✅ Günlük seviyesi filtrelemesi (yalnızca eşik değerinde veya üzerindeki mesajlar günlüğe kaydedilir)
- ✅ Ölümcül seviye günlükleri üzerinde otomatik program çıkışı
- ✅ Basit ve hafif API

---

## Günlük Seviyeleri

Günlük seviyeleri hangi mesajların görüntüleneceğini kontrol eder. Daha düşük seviyeler daha ayrıntılıdır:

| Seviye | Sabit | Değer | Amaç |
|--------|-------|-------|------|
| Debug | `DebugLevel` | 0 | Geliştirme için ayrıntılı tanı bilgileri |
| Info | `InfoLevel` | 1 | Genel bilgilendirme mesajları |
| Warn | `WarnLevel` | 2 | Potansiyel olarak zararlı durumlara ilişkin uyarı mesajları |
| Error | `ErrorLevel` | 3 | Ciddi sorunlara ilişkin hata mesajları |
| Fatal | `FatalLevel` | 4 | Program çıkışına neden olan ölümcül hatalar |

**Örnek**: Logger seviyesini `WarnLevel` olarak ayarlarsanız, yalnızca Warn, Error ve Fatal mesajları günlüğe kaydedilecektir. Debug ve Info mesajları yoksayılacaktır.

---

## Fonksiyonlar

### NewLogger

Belirtilen günlük seviyesi ile yeni bir Logger örneği oluşturur.

```go
func NewLogger(level LogLevel) *Logger
```

**Parametreler:**
- `level` - İlk günlükleme seviyesi eşiği

**Döner:**
- Yeni bir Logger örneğine işaretçi

**Örnek:**
```go
logger := gosugar.NewLogger(gosugar.InfoLevel)
```

---

### SetLevel

Logger'ın günlük seviyesi eşiğini çalışma zamanında günceller.

```go
func (l *Logger) SetLevel(level LogLevel)
```

**Parametreler:**
- `level` - Yeni günlükleme seviyesi eşiği

**Örnek:**
```go
logger.SetLevel(gosugar.DebugLevel) // Debug mesajlarını etkinleştir
```

---

### Debug

DEBUG seviyesinde bir mesaj günlüğe kaydeder.

```go
func (l *Logger) Debug(msg string)
```

**Parametreler:**
- `msg` - Günlüğe kaydedilecek mesaj

**Örnek:**
```go
logger.Debug("Veritabanı bağlantısı kuruldu")
```

---

### Info

INFO seviyesinde bir mesaj günlüğe kaydeder.

```go
func (l *Logger) Info(msg string)
```

**Parametreler:**
- `msg` - Günlüğe kaydedilecek mesaj

**Örnek:**
```go
logger.Info("Sunucu port 8080'de başlatıldı")
```

---

### Warn

WARN seviyesinde bir mesaj günlüğe kaydeder.

```go
func (l *Logger) Warn(msg string)
```

**Parametreler:**
- `msg` - Günlüğe kaydedilecek mesaj

**Örnek:**
```go
logger.Warn("Kullanımdan kaldırılan API uç noktası kullanıldı")
```

---

### Error

ERROR seviyesinde bir mesaj günlüğe kaydeder.

```go
func (l *Logger) Error(msg string)
```

**Parametreler:**
- `msg` - Günlüğe kaydedilecek mesaj

**Örnek:**
```go
logger.Error("Veritabanına bağlanma başarısız")
```

---

### Fatal

FATAL seviyesinde bir mesaj günlüğe kaydeder ve programı 1 koduyla kapatır.

```go
func (l *Logger) Fatal(msg string)
```

**Parametreler:**
- `msg` - Günlüğe kaydedilecek mesaj

**Örnek:**
```go
logger.Fatal("Kritik yapılandırma hatası - kapanıyor")
// Bu çağrıdan sonra program çıkış yapar
```

---

## Örnekler

### Temel Kurulum

```go
package main

import (
	"github.com/yourname/gosugar"
)

func main() {
	// Info seviyesinde logger oluştur
	logger := gosugar.NewLogger(gosugar.InfoLevel)

	logger.Debug("Bu gösterilmeyecek")       // Eşiğin altında
	logger.Info("Uygulama başlatıldı")       // Gösterilir: [2026-02-10 15:30:45] [INFO] Uygulama başlatıldı
	logger.Warn("Düşük bellek")               // Gösterilir: [2026-02-10 15:30:45] [WARN] Düşük bellek
}
```

### Çalışma Zamanında Günlük Seviyesini Değiştirme

```go
logger := gosugar.NewLogger(gosugar.WarnLevel)

logger.Info("Bu gösterilmeyecek")  // Eşiğin altında

logger.SetLevel(gosugar.InfoLevel)
logger.Info("Şimdi bu gösterilecek")  // Yeni eşikte
```

### Farklı Günlük Seviyelerini Kullanma

```go
logger := gosugar.NewLogger(gosugar.DebugLevel)

logger.Debug("Ayrıntılı tanı bilgileri")
logger.Info("Genel bilgiler")
logger.Warn("Bir şey hakkında uyarı")
logger.Error("Bir hata oluştu")
logger.Fatal("Kritik hata - çıkılıyor") // Program sonlandırılır
```

---

## Tasarım Kararları

### Seviye Tabanlı Filtreleme

Mesajlar ancak seviyesi logger'ın geçerli seviyesi ile eşitse veya daha yüksekse günlüğe kaydedilir. Bu, aşağıdakilere izin verir:
- Geliştirme sırasında `DebugLevel` ile her şeyi görmek
- Üretimde yalnızca kritik sorunları görmek için `ErrorLevel`'e geçmek

### Zaman Damgası Biçimlendirmesi

Her günlük mesajı `YYYY-AA-GG SS:DD:SS` biçiminde bir zaman damgası içerir. Bu, aşağıdakileri izlemeye yardımcı olur:
- Olaylar ne zaman meydana geldi
- Olaylar arasındaki süre
- Günlüklerde zamana dayalı desenler

### Ölümcül Programı Kapatır

`Fatal` yöntemi günlüğe kaydettikten sonra `os.Exit(1)` çağırır. Bu, aşağıdakileri sağlar:
- Kritik hatalar kapatılmadan önce günlüğe kaydedilir
- Program sıfır olmayan bir çıkış koduyla sonlandırılır (başarısızlığı gösterir)
- Çağıran kodun ölümcül hatalar için ayrı olarak işlenmesi gerekmez

### Yalnızca Dize API'si

Günlük yöntemleri yalnızca dizeler kabul eder, biçimlendirilmiş bağımsız değişkenler değil. Bu tasarım seçimi:
- API'yi basit ve hafif tutar
- Günlüğe kaydedilmeden önce dize oluşturma/biçimlendirmeyi teşvik eder
- Günlük çıktısını öngörülebilir ve tutarlı hale getirir

### Yalnızca Stdout

Tüm mesajlar standart çıktıya gider. Daha karmaşık günlükleme gereksinimleri için:
- Stdout'u kabukta bir dosyaya yönlendirin
- Bu logger'ın etrafında bir sarmalayıcı kullanın
- Daha gelişmiş günlükleme kütüphanelerini göz önünde bulundurun
