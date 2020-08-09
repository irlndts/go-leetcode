package leetcode

type Node struct {
	Val  int
	Left *Node
}

type Stack struct {
	Tail *Node
}

func (s *Stack) Push(val int) {
	tail := s.Tail
	s.Tail = &Node{
		Val:  val,
		Left: tail,
	}
}

func (s *Stack) Pop() int {
	val := s.Tail.Val
	s.Tail = s.Tail.Left
	return val
}

func (s *Stack) Peek() int {
	return s.Tail.Val
}

type MyQueue struct {
	Primary *Stack
	Second  *Stack

	Size int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		Primary: &Stack{},
		Second:  &Stack{},
		Size:    0,
	}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.Primary.Push(x)
	q.Size++
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {

	q.Size--

	if q.Second.Tail != nil {
		return q.Second.Pop()
	}

	for q.Primary.Tail != nil {
		q.Second.Push(q.Primary.Pop())
	}

	return q.Second.Pop()
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if q.Second.Tail != nil {
		return q.Second.Peek()
	}

	for q.Primary.Tail != nil {
		q.Second.Push(q.Primary.Pop())
	}

	return q.Second.Peek()
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return q.Size == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
