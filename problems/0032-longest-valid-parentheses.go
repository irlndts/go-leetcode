package problems

func longestValidParentheses(s string) int {

	result := 0
	stack := NewStack()
	stack.Add(-1)
	for i, r := range s {
		switch r {
		case '(':
			stack.Add(i)
		case ')':
			stack.Pop()
			if stack.Empty() {
				stack.Add(i)
			} else {
				result = max(result, i-stack.Peek())
			}
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Stack struct {
	Tail *Item
}

func (s *Stack) Add(val int) {
	if s.Tail == nil {
		item := &Item{Val: val}
		s.Tail = item
		return
	}

	s.Tail = &Item{Val: val, Left: s.Tail}
}

func (s *Stack) Pop() (int, bool) {
	if s.Tail == nil {
		return 0, true
	}

	item := s.Tail
	s.Tail = item.Left
	return item.Val, false
}

func (s *Stack) Empty() bool {
	return s.Tail == nil
}

func (s *Stack) Peek() int {
	return s.Tail.Val
}

type Item struct {
	Val  int
	Left *Item
}

func NewStack() *Stack {
	return &Stack{
		Tail: nil,
	}
}
