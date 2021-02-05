package main

import (
	"fmt"
	"learngo/demo"
)

func main() {
	l := demo.NewLinkedList()
	_ = l.Add(1, 2, 3)
	l1 := demo.NewLinkedList()
	_ = l1.Add(1, 2, 2)
	fmt.Println(l.Equals(l1))
}
