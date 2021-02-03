package demo

import (
	"errors"
	"fmt"
)

type Node struct {
	Data  interface{}
	Prior *Node
	Next  *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) InitLinkedList() {
	l.head = new(Node)
	l.head.Next = l.head
	l.head.Prior = l.head
	l.length = 0
}

// 向list末尾加入元素
func (l *LinkedList) Add(obj ...interface{}) error {
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

func (l *LinkedList) Insert(location int, obj interface{}) error {
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

func (l *LinkedList) Set(location int, obj interface{}) error {
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

func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) Get(location int) (interface{}, error) {
	if location < 0 {
		return nil, errors.New("Index should be greater than 0 ")
	} else if location > l.length {
		return nil, errors.New("Index out of range ")
	} else {
		cur := l.head
		for i := 0; i < (location - 1); i++ {
			cur = cur.Next
		}
		return cur.Data, nil
	}
}

func (l *LinkedList) Equals(list List) bool {
	flag := true
	cur := l.head
	i := 1
	var j interface{}
	for cur != nil {
		j, _ = list.Get(i)
		i++
		if cur.Data != j {
			return false
		} else {
			flag = true
			cur = cur.Next
		}
		if i > l.length {
			return false
		}

	}
	return flag
}

func (l *LinkedList) ToSlice() []interface{} {
	cur := l.head
	s := make([]interface{}, 1)
	for cur != nil {
		s = append(s, cur.Data)
		cur = cur.Next
	}
	return s
}

func (l *LinkedList) Size() int {
	return l.length
}

func (l *LinkedList) Print() {
	cur := l.head
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
}
