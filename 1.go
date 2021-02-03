package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"unicode"
)

// 用来统计字符个数
type count struct {
	HanCount   int
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

var instance *count

// Once能保证某个操作仅且只执行一次
var once sync.Once

// 实现单例
func GetCount() *count {
	once.Do(func() {
		instance = &count{}
	})
	return instance
}

func statistics(s string) *count {
	a := GetCount()
	for _, v := range s {
		// 调用unicode的方法来判断是不是汉字
		if unicode.Is(unicode.Han, v) {
			a.HanCount++
		} else if 'A' <= v && v <= 'Z' || 'a' <= v && v <= 'z' {
			a.ChCount++
		} else if '1' <= v && v <= '9' {
			a.NumCount++
		} else if v == 32 {
			a.SpaceCount++
		} else {
			a.OtherCount++
		}
	}
	return a
}

func main() {
	fmt.Println("请输入一行字符串")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	c := statistics(text)
	fmt.Println("汉字的个数：", c.HanCount)
	fmt.Println("英文的个数：", c.ChCount)
	fmt.Println("数字的个数：", c.NumCount)
	fmt.Println("空格的个数：", c.SpaceCount)
	fmt.Println("其他字符的个数：", c.OtherCount)
}
