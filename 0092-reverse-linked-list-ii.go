package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	index := 1
	start := head
	var prev *ListNode
	for head != nil && index != m {
		prev = head
		head = head.Next
		index++
	}
	head = prev
	if m == 1 {
		start = reverse(start, n-index)
	} else {
		head.Next = reverse(head.Next, n-index)
	}
	return start
}

func reverse(head *ListNode, n int) *ListNode {
	var cur, start *ListNode
	last := head
	index := 0
	for head != nil && index <= n {
		start = head
		head = head.Next
		start.Next = cur
		cur = start
		index++
	}

	last.Next = head
	return start
}

func p(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
