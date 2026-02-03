package gosugar

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//
// RANDOM INT (min inclusive, max inclusive)
//

func RandInt(min, max int) int {
	if min > max {
		panic("min cannot be greater than max")
	}
	return rand.Intn(max-min+1) + min
}

//
// RANDOM FLOAT (min inclusive, max exclusive)
//

func RandFloat(min, max float64) float64 {
	if min >= max {
		panic("min must be less than max")
	}
	return min + rand.Float64()*(max-min)
}

//
// RANDOM BOOL
//

func RandBool() bool {
	return rand.Intn(2) == 1
}

//
// CHOICE (pick random element)
//

func Choice[T any](items []T) T {
	if len(items) == 0 {
		panic("cannot choose from empty slice")
	}
	return items[rand.Intn(len(items))]
}

//
// RANDOM STRING (letters only)
//

func RandString(length int) string {
	if length <= 0 {
		panic("length must be positive")
	}

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
