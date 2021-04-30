package problems

import "fmt"

// https://leetcode.com/problems/linked-list-cycle-ii/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var cycle bool
	slow, fast := head, head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		fmt.Println(slow.Val, fast.Val)

		if slow == fast {
			cycle = true
			break
		}
	}

	if !cycle {
		return nil
	}

	start := head
	for start != fast {
		start = start.Next
		fast = fast.Next
	}

	return start
}
