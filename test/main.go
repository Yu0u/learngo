package main

import (
	"bufio"
	"fmt"
	"learngo/calculator"
	"os"
)

func main() {
	for {
		fmt.Println("请输入要计算的表达式：(按Q退出)")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "Q" {
			fmt.Println("退出陈坤")
			break
		}
		calculate := calculator.Calculate(text)
		fmt.Println("结果是 ", calculate)
	}
}
