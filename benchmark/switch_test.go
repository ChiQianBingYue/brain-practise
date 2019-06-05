package benchmark

import "testing"

func BenchmarkSwitchWithTemp(b *testing.B) {
	// run the Fib function b.N times

	for n := 0; n < b.N; n++ {
		a, b := 1, 2
		temp := a
		a = b
		b = temp
	}
}

func BenchmarkSwitchWithArray(b *testing.B) {

	for n := 0; n < b.N; n++ {
		a, b := 1, 2
		a, b = b, a
	}
}
