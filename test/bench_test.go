package main

import (
	"testing"

	"github.com/coderianx/gosugar"
)

// Benchmark: RandInt
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gosugar.RandInt(1, 100)
	}
}

// Benchmark: RandString
func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gosugar.RandString(32)
	}
}

// Benchmark: RandBool
func BenchmarkRandBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gosugar.RandBool()
	}
}

// Benchmark: RandFloat
func BenchmarkRandFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gosugar.RandFloat(0.0, 100.0)
	}
}

// Benchmark: Choice
func BenchmarkChoice(b *testing.B) {
	items := []int{1, 2, 3, 4, 5, 10, 20, 30}
	for i := 0; i < b.N; i++ {
		gosugar.Choice(items)
	}
}

// Benchmark: EnvString
func BenchmarkEnvString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gosugar.EnvString("PATH", "default")
	}
}

// Benchmark: ReadFile
func BenchmarkReadFile(b *testing.B) {
	// Test dosyası oluştur
	gosugar.WriteFile("/tmp/benchmark_test.txt", "Hello, World!")
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gosugar.ReadFile("/tmp/benchmark_test.txt")
	}
}

// Benchmark: WriteFile
func BenchmarkWriteFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gosugar.WriteFile("/tmp/benchmark_write.txt", "Test content")
	}
}

// Benchmark: Must
func BenchmarkMust(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = gosugar.Must(42, nil)
	}
}

// Benchmark: Try
func BenchmarkTry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = gosugar.Try(func() int {
			return 42
		})
	}
}
