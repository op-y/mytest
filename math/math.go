package math

import (
	"github.com/op-y/mytest/function"
)

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func Rand10() int {
	// 返回10以内的整数
	return function.RandInt(10)
}
