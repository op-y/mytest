package function_test

import (
	"testing"

	"github.com/op-y/mytest/function"
	"github.com/op-y/mytest/math"
)

func TestFibonacci(t *testing.T) {
	fibs := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	idx := math.Rand10()
	expected := fibs[idx]
	result := function.Fibonacci(idx)
	if result != expected {
		t.Fatalf("the %d fibonacci number is %d, but expect %d", idx, result, expected)
	}
	t.Logf("the %d fibonacci number is %d, test pass", idx, result)
}

func BenchmarkFibonacci1h(b *testing.B) {
	for range b.N {
		num := 10
		function.Fibonacci(num)
	}
}

func FuzzFibonacci(f *testing.F) {
	for _, seed := range [][]int{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
		{9, 55},
	} {
		f.Add(seed[0], seed[1])
	}
	f.Fuzz(func(t *testing.T, in, expected int) {
		result := function.Fibonacci(in)
		if result != expected {
			t.Fatalf("the %d fibonacci number is %d, but expect %d", in, result, expected)
		}
	})
}
