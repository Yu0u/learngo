package demo

import (
	"errors"
	"fmt"
	"sync"
)

type ArrayList struct {
	data []interface{}
	size int
	lock sync.Mutex
}

type ArrayListIterator struct {
	list         *ArrayList
	currentIndex int
}

func (a *ArrayList) Iterator() Iterator {
	iterator := new(ArrayListIterator)
	iterator.currentIndex = 0
	iterator.list = a
	return iterator
}

func (a *ArrayListIterator) HasNext() bool {
	return a.currentIndex < a.list.Size()
}

func (a *ArrayListIterator) Next() (interface{}, error) {
	if !a.HasNext() {
		return nil, errors.New("Not have next element ")
	}
	value, err := a.list.Get(a.currentIndex)
	a.currentIndex++
	return value, err
}

func (a *ArrayListIterator) getIndex() int {
	return a.currentIndex
}

func NewArrayList() *ArrayList {
	// 分配一个容量为0的切片，
	// 如果切片容量不为零，则会在切片初始分配后添加
	return &ArrayList{
		data: make([]interface{}, 0),
		size: 0,
	}
}

func (a *ArrayList) Add(obj ...interface{}) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	for _, v := range obj {
		a.data = append(a.data, v)
		a.size++
	}
	return nil
}

func (a *ArrayList) Insert(location int, obj interface{}) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	if location < 0 || location >= a.size {
		return errors.New("Index out of range ")
	}
	// 将从location开始的元素存到一个辅助切片中
	rear := append([]interface{}{}, a.data[location:]...)
	// 先将需要插入的元素插入到location之前，然后再将rear切片插入到要插入元素的后面
	a.data = append(append(a.data[:location], obj), rear...)
	a.size++
	return nil
}

func (a *ArrayList) Set(location int, obj interface{}) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	if location < 0 || location >= a.size {
		return errors.New("Index out of range ")
	}
	// 根据索引修改值
	a.data[location] = obj
	return nil
}

func (a *ArrayList) Contain(obj interface{}) bool {
	for _, v := range a.data {
		// 遍历是否包含
		if v == obj {
			return true
		}
	}
	return false
}

func (a *ArrayList) IsEmpty() bool {
	// 如果 size 为 0 则 ArrayList 为空
	return a.size == 0
}

func (a *ArrayList) Get(location int) (interface{}, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if location < 0 || location >= a.size {
		return nil, errors.New("Index out of range ")
	}
	// 根据索引查找值
	i := a.data[location]
	return i, nil
}

func (a *ArrayList) Equals(list List) bool {
	flag := false
	iterator1 := list.Iterator()
	iterator2 := a.Iterator()
	for iterator1.HasNext() && iterator2.HasNext() {
		o1, _ := iterator1.Next()
		o2, _ := iterator2.Next()
		if o1 == o2 {
			flag = true
		}
	}
	flag = !(iterator1.HasNext() || iterator2.HasNext())
	return flag
}

func (a *ArrayList) ToSlice() []interface{} {
	// 因为data已经是一个slice 所以可以直接返回
	r := a.data
	return r
}

func (a *ArrayList) Size() int {
	return a.size
}

// 测试用
func (a *ArrayList) Print() {
	for i := 0; i < a.size; i++ {
		fmt.Print(a.data[i], " ")
	}
	fmt.Println()
}
