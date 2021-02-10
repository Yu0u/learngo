package main

import (
	"fmt"
	"learngo/demo"
)

func main() {
	l := demo.NewLinkedList()

	_ = l.Add(1, 2, 3)

	l2 := demo.NewLinkedList()
	_ = l2.Add(1, 2, 3)

	fmt.Println(l.Equals(l2))
}
