package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if k == 0 {
		return head
	}

	start := head

	length := 0
	for head != nil {
		length++
		head = head.Next
	}

	k = k % length
	if k == 0 {
		return start
	}

	i := 0
	head = start
	l, r := head, head
	for r.Next.Next != nil {
		head = head.Next
		i++
		if i >= k {
			l = l.Next
		}
		r = r.Next
	}

	r.Next.Next = start
	next := l.Next
	l.Next = nil

	return next
}
