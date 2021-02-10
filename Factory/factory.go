package main

import "fmt"

type animal interface {
	Call()
}

type Cat struct {
}

func (c *Cat) Call() {
	fmt.Println("喵喵喵")
}

type Dog struct {
}

func (d *Dog) Call() {
	fmt.Println("汪汪汪")
}

type AnimalFactory interface {
	Create() animal
}

type CatFactory struct {
}

func (c *CatFactory) Create() animal {
	return &Cat{}
}

type DogFactory struct {
}

func (d *DogFactory) Create() animal {
	return &Dog{}
}

func CreateCatFactory() AnimalFactory {
	return &CatFactory{}
}

func CreateDogFactory() AnimalFactory {
	return &DogFactory{}
}

func main() {
	dog := CreateDogFactory()
	dog.Create().Call()

	cat := CreateCatFactory()
	cat.Create().Call()
}
