package demo

import (
	"errors"
	"fmt"
	"sync"
)

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	head   *Node
	length int
	lock   sync.Mutex
}

type LinkedListIterator struct {
	list         *LinkedList
	currentIndex int
}

func (l *LinkedListIterator) HasNext() bool {
	return l.currentIndex < l.list.Size()
}

func (l *LinkedListIterator) Next() (interface{}, error) {
	if !l.HasNext() {
		return nil, errors.New("Not have next element ")
	}
	value, err := l.list.Get(l.currentIndex)
	l.currentIndex++
	return value, err
}

func (l *LinkedListIterator) getIndex() int {
	return l.currentIndex
}

func (l *LinkedList) Iterator() Iterator {
	iterator := new(LinkedListIterator)
	iterator.currentIndex = 1
	iterator.list = l
	return iterator
}

func NewLinkedList() *LinkedList {
	return &LinkedList{length: 0}
}

// 向list末尾加入元素
func (l *LinkedList) Add(obj ...interface{}) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	for i := 0; i < len(obj); i++ {
		node := &Node{Data: obj[i]}
		if l.IsEmpty() {
			l.head = node
		} else {
			cur := l.head
			for cur.Next != nil {
				cur = cur.Next
			}
			cur.Next = node
		}
		l.length++
	}
	return nil
}

func (l *LinkedList) AddFromHead(obj ...interface{}) *Node {
	l.lock.Lock()
	defer l.lock.Unlock()
	var node *Node
	for i := 0; i < len(obj); i++ {
		node = &Node{Data: obj[i]}
		if l.IsEmpty() {
			l.head = node

		}
		node.Next = l.head
		l.head = node
		l.length++
	}
	return node
}

// 通过遍历找到location，并插入
func (l *LinkedList) Insert(location int, obj interface{}) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if location < 0 {
		return errors.New("Index should be greater than 0 ")
	} else if location <= 1 {
		l.AddFromHead(obj)
	} else if location > l.length+1 {
		return errors.New("Index out of range ")
	} else {
		pre := l.head
		count := 1

		for count < (location - 1) {
			pre = pre.Next
			count++
		}
		node := &Node{Data: obj}
		node.Next = pre.Next
		pre.Next = node
		l.length++
	}
	return nil
}

// 通过遍历找到值，并修改
func (l *LinkedList) Set(location int, obj interface{}) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if location < 0 {
		return errors.New("Index should be greater than 0 ")
	} else if location > l.length+1 {
		return errors.New("Index out of range ")
	} else {
		cur := l.head
		for i := 0; i < (location - 1); i++ {
			cur = cur.Next
		}
		cur.Data = obj
		return nil
	}
}

func (l *LinkedList) Contain(obj interface{}) bool {
	cur := l.head
	for cur != nil {
		if cur.Data == obj {
			return true
		}
		cur = cur.Next
	}
	return false
}

// 头节点为空则为空
func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

// 遍历查值
func (l *LinkedList) Get(location int) (interface{}, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if location < 0 {
		return nil, errors.New("Index should be greater than 0 ")
	} else if location > l.length {
		return nil, errors.New("Index out of range ")
	} else {
		cur := l.head
		for i := 0; i < (location - 1); i++ {
			cur = cur.Next
		}
		data := cur.Data
		return data, nil
	}
}

//
func (l *LinkedList) Equals(list List) bool {
	flag := false
	iterator1 := list.Iterator()
	iterator2 := l.Iterator()
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

// 通过遍历将数据转换为slice
func (l *LinkedList) ToSlice() []interface{} {
	cur := l.head
	s := make([]interface{}, 0)
	for cur != nil {
		s = append(s, cur.Data)
		cur = cur.Next
	}
	return s
}

func (l *LinkedList) Size() int {
	return l.length
}

// 测试用
func (l *LinkedList) Print() {
	cur := l.head
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
}
