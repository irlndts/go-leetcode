package leetcode

// https://leetcode.com/problems/min-stack/

type Node struct {
	Val  int
	Left *Node
}

type MinStack struct {
	Tail *Node
	Mins *Node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	tail := this.Tail
	this.Tail = &Node{
		Val:  x,
		Left: tail,
	}

	mins := this.Mins
	val := x
	if mins != nil && mins.Val < x {
		val = mins.Val
	}

	this.Mins = &Node{
		Val:  val,
		Left: mins,
	}
}

func (this *MinStack) Pop() {
	if this.Tail == nil || this.Tail.Left == nil {
		this.Tail = nil
		this.Mins = nil
		return
	}

	this.Tail = this.Tail.Left
	this.Mins = this.Mins.Left
}

func (this *MinStack) Top() int {
	return this.Tail.Val
}

func (this *MinStack) GetMin() int {
	return this.Mins.Val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
