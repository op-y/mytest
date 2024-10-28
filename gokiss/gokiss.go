package gokiss

import (
	"fmt"

	"github.com/go-kiss/monkey"
)

func sum(a, b int) int { return a + b }

func sub1[T int | float64](a, b T) T { return a - b }
func sub2[T int | float64](a, b T) T { return a + b }

type S struct {
	i int
}

func (s *S) Get() int {
	return s.i
}

func RunGoKissMock() {
	// mock 普通函数
	monkey.Patch(sum, func(a, b int) int { return a - b })
	fmt.Println(sum(1, 2)) // 输出 -1

	// mock 泛型函数
	monkey.Patch(sub1[int], sub2[int])
	fmt.Println(sub1(1, 2)) // 输出 3

	// mock 结构体方法
	monkey.Patch((*S).Get, func(s *S) int { return -1 })
	var s S
	fmt.Println(s.Get()) // 输出 -1
}
