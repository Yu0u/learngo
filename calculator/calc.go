package calculator

import (
	"bytes"
	"learngo/stack"
	"strconv"
)

type operation struct {
	op       byte
	priority int
}

func Calculate(s string) int {
	var b bytes.Buffer

	for i := 0; i < len(s); i++ {
		if (s[i] == '-' || s[i] == '+') && (i == 0 || s[i-1] == '(') {
			b.WriteByte('0')
		}
		b.WriteByte(s[i])
	}
	b.WriteByte('+')
	str := b.String()
	numStack := stack.Construct()
	opStack := stack.Construct()

	hash := make(map[byte]int, 0)
	hash['+'] = 1
	hash['-'] = 1
	hash['*'] = 2
	hash['/'] = 2

	prio := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}
		if _, ok := hash[str[i]]; ok {
			oper := operation{str[i], hash[str[i]] + prio}
			for !opStack.Empty() && opStack.Peek().(operation).priority >= oper.priority {
				op := opStack.Pop().(operation).op
				cur := 0
				b := numStack.Pop().(int)
				a := numStack.Pop().(int)
				switch op {
				case '+':
					cur = a + b
				case '-':
					cur = a - b
				case '*':
					cur = a * b
				case '/':
					cur = a / b
				}
				numStack.Push(cur)
			}
			opStack.Push(oper)
		} else if str[i] == '(' {
			prio += 10
		} else if str[i] == ')' {
			prio -= 10
		} else {
			tmp := string(str[i])
			for str[i+1] >= '0' && str[i+1] <= '9' {
				tmp += string(str[i+1])
				i++
			}
			num, _ := strconv.Atoi(tmp)
			numStack.Push(num)
		}
	}
	return numStack.Pop().(int)
}
