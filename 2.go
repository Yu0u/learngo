package main

import "fmt"

// 定义一个fib函数，函数返回值为匿名函数
func fib() func(int) int {
	a, b := 0, 1
	return func(i int) int {
		// 从0开始输出值，没有这个if会从1开始输出
		if i > 0 {
			a, b = b, a+b
		}
		return a
	}
}

func main() {
	f := fib()
	for i := 0; i < 50; i++ {
		// 传递i过去只是为了输出0
		fmt.Print(f(i), " ")
	}
}
