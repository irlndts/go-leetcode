package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	stack := NewStack()

	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			a, b := stack.Pop(), stack.Pop()
			switch v {
			case "+":
				stack.Push(b + a)
			case "-":
				stack.Push(b - a)
			case "*":
				stack.Push(b * a)
			case "/":
				stack.Push(b / a)
			}
		default:
			val, _ := strconv.Atoi(v)
			stack.Push(val)
		}
	}
	return stack.Pop()
}

// Node is an element in stack.
type Node struct {
	Val  int
	Left *Node
}

// Stack is a stack instance.
type Stack struct {
	Tail *Node
}

// NewStack returns new stack instance.
func NewStack() *Stack {
	return &Stack{}
}

// Push adds an element to stack.
func (s *Stack) Push(val int) {
	tail := s.Tail
	s.Tail = &Node{
		Val:  val,
		Left: tail,
	}
}

// Pop returns an element from stack.
func (s *Stack) Pop() int {
	val := s.Tail.Val
	s.Tail = s.Tail.Left
	return val
}
