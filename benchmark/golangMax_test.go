package benchmark

import (
	"math"
	"testing"
)

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func mathMath(i, j int) int {
	return int(math.Min(float64(i), float64(j)))
}

func BenchmarkMax(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		max(12321, 3210)
	}
}
func BenchmarkMathMax(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		mathMath(12321, 3210)
	}
}
