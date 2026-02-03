package gosugar

//
// MUST
// value, err dönen fonksiyonlar için
//

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

//
// CHECK
// sadece err dönen fonksiyonlar için
//

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

//
// TRY
// panic olabilecek kodu güvenli çalıştırır
//

func Try[T any](fn func() T) (v T, ok bool) {
	defer func() {
		if recover() != nil {
			ok = false
		}
	}()
	return fn(), true
}

//
// OR
// Try ile birlikte default fallback
//

func Or[T any](v T, ok bool, fallback T) T {
	if ok {
		return v
	}
	return fallback
}

//
// IGNORE
// hatayı bilinçli şekilde yutmak için
//

func Ignore(err error) {
	_ = err
}
