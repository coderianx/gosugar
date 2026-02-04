# GoSugar Dokumentasyon

Hoş geldiniz! Bu dokümantasyon GoSugar kütüphanesinin derinlemesine açıklamasıdır. Projeyi hiç tanımayan bir geliştirici için başlamaktan bitirmeye kadar her şeyi bulacaksınız.

## 📚 Dokümantasyonu Nasıl Kullanacaksınız?

### **1. Başlangıç Seviyesi (Hızlı Giriş)**
Eğer projeyi ilk defa görüyorsanız:
- [`getting-started.md`](guides/getting-started.md) ile başlayın
- Projenin ne yaptığını, temel kavramları öğrenin
- Basit örneklerle çalıştırın

### **2. Mimarisi Anlamak İçin**
Projenin iç yapısını, tasarım kararlarını, veri akışını merak ediyorsanız:
- [`ARCHITECTURE.md`](architecture/ARCHITECTURE.md) okuyun
- Neden böyle yapıldığını anlayın
- Modüller arası ilişkiler grafikle görün

### **3. API Referansı**
Belirli bir fonksiyonun nasıl çalıştığını bilmek istiyorsanız:
- [`api/`](api/) klasöründe modül başlıklı dosyalar var
- Her modülün tüm fonksiyonları, parametreleri, dönüş değerleri
- Örnekler ve taşla testilmiş use-case'ler

**Modül Referansları:**
- [`env.md`](api/env.md) - Ortam değişkenleri yönetimi
- [`input.md`](api/input.md) - Kullanıcıdan input alma
- [`validators.md`](api/validators.md) - Input doğrulama
- [`random.md`](api/random.md) - Rastgele veri üretimi
- [`errors.md`](api/errors.md) - Hata yönetimi
- [`file.md`](api/file.md) - Dosya işlemleri
- [`http.md`](api/http.md) - HTTP istek işlemleri

### **4. Rehberler & Öğretim**
Özel senaryolarda nasıl kullanılacağını öğrenmek için:
- [`guides/`](guides/) klasöründeki rehberleri okuyun
- CLI uygulaması yazma, hata yönetimi, best practice'ler

**Mevcut Rehberler:**
- [`getting-started.md`](guides/getting-started.md) - Başlama kılavuzu
- [`design-patterns.md`](guides/design-patterns.md) - Tasarım desenleri
- [`error-handling.md`](guides/error-handling.md) - Hata yönetimi
- [`testing-with-gosugar.md`](guides/testing-with-gosugar.md) - Test yazma

### **5. Tasarım & İç Yapı**
Projeyi geliştirmek veya katkıda bulunmak istiyorsanız:
- [`architecture/ARCHITECTURE.md`](architecture/ARCHITECTURE.md) - Tam mimari
- [`architecture/design-decisions.md`](architecture/design-decisions.md) - Tasarım kararları

## 🎯 Hızlı Navigasyon

| Sorunuz | Gidecek Yer |
|---------|-----------|
| "GoSugar nedir?" | [`getting-started.md`](guides/getting-started.md) |
| "`EnvString()` nasıl çalışır?" | [`api/env.md`](api/env.md) |
| "Hata nasıl ele alırım?" | [`error-handling.md`](guides/error-handling.md) / [`api/errors.md`](api/errors.md) |
| "CLI uygulaması yazabilir miyim?" | [`getting-started.md`](guides/getting-started.md) + [`api/input.md`](api/input.md) |
| "Mimarisi nasıl?" | [`ARCHITECTURE.md`](architecture/ARCHITECTURE.md) |
| "Test yazabilir miyim?" | [`testing-with-gosugar.md`](guides/testing-with-gosugar.md) |
| "Neden panic kullanılıyor?" | [`design-decisions.md`](architecture/design-decisions.md) |
| "Tüm fonksiyonların listesi?" | API referansında modül başlığı seçin |

## 🗂️ Dokümantasyon Yapısı

```
docs/
├── README.md (⬅️ Siz burasınız)
├── guides/
│   ├── getting-started.md          # İlk adımlar
│   ├── design-patterns.md          # Tasarım desenleri
│   ├── error-handling.md           # Hata yönetimi stratejileri
│   └── testing-with-gosugar.md    # Test yazma
├── api/
│   ├── env.md                      # Ortam değişkenleri
│   ├── input.md                    # Kullanıcı inputu
│   ├── validators.md               # Validatörler
│   ├── random.md                   # Rastgele veri
│   ├── errors.md                   # Hata yönetimi
│   ├── file.md                     # Dosya işlemleri
│   └── http.md                     # HTTP istemci
└── architecture/
    ├── ARCHITECTURE.md             # Tam mimari açıklama
    └── design-decisions.md         # Tasarım kararları açıklı
```

## ⏱️ Tahmini Okuma Süreleri

- **Hızlı başlangıç**: ~15 dakika (`getting-started.md`)
- **Bir modülü derinlemesine anlamak**: ~10-15 dakika (API referansı)
- **Mimariyiyi tamamen anlamak**: ~30-40 dakika (`ARCHITECTURE.md`)
- **Tüm dokümantasyonu okumak**: ~2 saat

## 💡 Önemli Varsayımlar

Bu dokümantasyon aşağıdakileri varsayar:
- ✅ Go 1.18+ hakkında temel bilgi (paketler, fonksiyonlar, generics)
- ✅ Terminal/CLI uygulamaları hakkında bilgi
- ❌ GoSugar hakkında önceki bilgi **gerekli değil**

## 🤝 Katkı & Sorular

Dokümantasyonu iyileştirmek için:
1. Eksik gördüğünüz konuları belirtebilirsiniz
2. Kötü açıklanan kısımlar için issue açabilirsiniz
3. Pull request ile dokümantasyon iyileştirmesi yapabilirsiniz

## 🔗 Hızlı Linkler

- **Ana Repository**: `github.com/coderianx/gosugar`
- **README.md**: Proje temel özeti (kütüphane kurulumu, lisanslama)
- **ROADMAP.md**: Planlanan yeni özellikler
- **info.md**: Eski/iç dokümantasyon (referans için)

---

**Not:** Kütüphane, modüler bir yapıya sahiptir. Her modül bağımsız olarak çalışabilir ve sadece gerekli olanı kullanabilirsiniz. Başlamaya hazır mısınız? → [`getting-started.md`](guides/getting-started.md) 🚀
