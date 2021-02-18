package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
	Name string
}

func (c *Cat) Sleep() {
	fmt.Println("Sleep...")
}

func (c *Cat) A() {
	fmt.Println("A...")
}

func (c *Cat) Eat() {
	fmt.Println("Eat...")
}

func main() {
	cat := Cat{}
	value := reflect.ValueOf(cat)
	f := value.Field(0)
	f.SetString("1")
	fmt.Printf("%#v", cat)
}
