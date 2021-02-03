package demo

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	data []interface{}
	size int
}

func (a *ArrayList) InitArrayList() {
	a.data = make([]interface{}, 4)
	a.size = 0
}

func (a *ArrayList) Add(obj ...interface{}) error {
	for _, v := range obj {
		a.data = append(a.data, v)
		a.size++
	}
	return nil
}

func (a *ArrayList) Insert(location int, obj interface{}) error {
	if location < 0 || location >= a.size {
		return errors.New("Index out of range ")
	}
	rear := append([]interface{}{}, a.data[location:]...)
	a.data = append(append(a.data[:location], obj), rear...)
	a.size++
	return nil
}

func (a *ArrayList) Set(location int, obj interface{}) error {
	if location < 0 || location >= a.size {
		return errors.New("Index out of range ")
	}
	a.data[location] = obj
	return nil
}

func (a *ArrayList) Contain(obj interface{}) bool {
	for _, v := range a.data {
		if v == obj {
			return true
		}
	}
	return false
}

func (a *ArrayList) IsEmpty() bool {
	return a.size == 0
}

func (a *ArrayList) Get(location int) (interface{}, error) {
	if location < 0 || location >= a.size {
		return nil, errors.New("Index out of range ")
	}
	return a.data[location], nil
}

func (a *ArrayList) Equals(list List) bool {
	flag := true
	var j interface{}
	for i := 0; i < a.size; i++ {
		j, _ = list.Get(i)
		if a.data[i] != j {
			return false
		} else {
			flag = true
		}
	}
	if list.Size() > a.Size() {
		flag = false
	}
	return flag
}

func (a *ArrayList) ToSlice() []interface{} {
	r := a.data
	return r
}

func (a *ArrayList) Size() int {
	return a.size
}

func (a *ArrayList) Print() {
	for i := 0; i < a.size; i++ {
		fmt.Print(a.data[i], " ")
	}
	fmt.Println()
}
