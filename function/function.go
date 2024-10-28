package function

import (
	"math/rand"
)

func RandInt(upper int) int {
	// 返回upper以内随机整数
	return rand.Intn(upper)
}

func Fibonacci(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return Fibonacci(num-1) + Fibonacci(num-2)
}
