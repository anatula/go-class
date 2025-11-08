package addbenchmark

import (
	"testing"
)

// Function to benchmark
func Add(a, b int) int {
	return a + b
}

// Benchmark function (must start with "Benchmark")
func BenchmarkAdd(b *testing.B) {
	// b.N is the number of iterations determined by the benchmark framework
	for i := 0; i < b.N; i++ {
		Add(i, i+1) // Use changing inputs to prevent optimization
	}
}
