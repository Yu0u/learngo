package stack

type Node struct {
	Val  interface{}
	Next *Node
}

type stack struct {
	dummy *Node
}

func Construct() *stack {
	return &stack{&Node{}}
}

func (s *stack) Push(x interface{}) {
	node := &Node{x, nil}
	node.Next = s.dummy.Next
	s.dummy.Next = node
}

func (s *stack) Pop() interface{} {
	val := s.dummy.Next.Val
	s.dummy.Next = s.dummy.Next.Next
	return val
}

func (s *stack) Peek() interface{} {
	return s.dummy.Next.Val
}

func (s *stack) Empty() bool {
	return s.dummy.Next == nil
}
