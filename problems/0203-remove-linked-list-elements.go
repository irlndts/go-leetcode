package problems

// https://leetcode.com/problems/remove-linked-list-elements/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return head
	}

	start := head

	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
			continue
		}

		head = head.Next
	}

	return start
}
