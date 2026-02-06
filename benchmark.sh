#!/bin/bash

# GoSugar Benchmark Script
# Test project performance

set -e

echo "================================"
echo "   GoSugar Benchmark Suite"
echo "================================"
echo ""

# Color codes
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Test directory
TEST_DIR="./test"
BENCHMARK_DIR="./benchmark"

# Create benchmark directory if it doesn't exist
mkdir -p "$BENCHMARK_DIR"

echo -e "${BLUE}1. Running Go Benchmark Tests...${NC}"
echo ""

# Run benchmarks if bench_test.go exists
if [ -f "${TEST_DIR}/bench_test.go" ]; then
    echo "Running benchmarks from test directory..."
    go test -bench=. -benchmem -benchtime=3s -run=^$ ./test
else
    echo -e "${YELLOW}⚠️  Benchmark test file not found${NC}"
    echo "Creating: ${TEST_DIR}/bench_test.go"
    
    # Create basic benchmark test file
    cat > "${TEST_DIR}/bench_test.go" << 'EOF'
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
EOF
    
    echo -e "${GREEN}✓ Benchmark test file created${NC}"
    echo ""
    echo "Running benchmarks..."
    go test -bench=. -benchmem -benchtime=3s -run=^$ ./test
fi

echo ""
echo -e "${BLUE}2. Memory Profiling${NC}"
echo ""

# Run tests with CPU and memory profiling
CPU_PROFILE="$BENCHMARK_DIR/cpu.prof"
MEM_PROFILE="$BENCHMARK_DIR/mem.prof"

echo "Writing CPU profile: $CPU_PROFILE"
go test -cpuprofile="$CPU_PROFILE" -bench=Rand -benchtime=2s -run=^$ ./test > /dev/null 2>&1

echo "Writing Memory profile: $MEM_PROFILE"
go test -memprofile="$MEM_PROFILE" -bench=Rand -benchtime=2s -run=^$ ./test > /dev/null 2>&1

echo -e "${GREEN}✓ Profiling completed${NC}"
echo ""

echo -e "${BLUE}3. Build Speed${NC}"
echo ""

# Measure build time
BUILD_START=$(date +%s%N)
go build -o /tmp/gosugar_test ./test/main.go 2>/dev/null
BUILD_END=$(date +%s%N)

BUILD_TIME=$(( (BUILD_END - BUILD_START) / 1000000 ))
echo -e "${GREEN}Build time: ${BUILD_TIME}ms${NC}"
echo ""

echo -e "${BLUE}4. Binary Size${NC}"
echo ""

# Measure binary size
BINARY_SIZE=$(ls -lh /tmp/gosugar_test | awk '{print $5}')
echo -e "${GREEN}Binary size: $BINARY_SIZE${NC}"
echo ""

echo -e "${BLUE}5. Test Coverage${NC}"
echo ""

# Generate coverage report
go test -cover -coverprofile="$BENCHMARK_DIR/coverage.out" ./test > /dev/null 2>&1

# Get coverage percentage
if command -v go &> /dev/null; then
    COVERAGE=$(go tool cover -func="$BENCHMARK_DIR/coverage.out" | tail -1 | awk '{print $NF}')
    echo -e "${GREEN}Test Coverage: $COVERAGE${NC}"
fi
echo ""

echo -e "${BLUE}6. Benchmark Results Summary${NC}"
echo ""

cat << 'EOF'
Benchmark profiles can be viewed with the following commands:

📊 CPU Profiling:
   go tool pprof benchmark/cpu.prof

💾 Memory Profiling:
   go tool pprof benchmark/mem.prof

📈 Coverage Report:
   go tool cover -html=benchmark/coverage.out

EOF

echo -e "${GREEN}✓ Benchmark completed!${NC}"
echo ""

# Profiling tips
cat << 'EOF'
🎯 Profiling Types:

1. **CPU Profile** - Which functions use the most CPU?
2. **Memory Profile** - Which functions allocate the most memory?
3. **Coverage** - Which code did we test?

💡 Tips:
- To compare benchmark results: benchstat
  go install golang.org/x/perf/cmd/benchstat@latest
  
- A good benchmark should run for 3+ seconds (benchtime=3s)
- Smaller ns/op (nanosecond per operation) means faster performance

EOF

echo "================================"
echo "   Benchmark Completed"
echo "================================"
