package leetcode

// https://leetcode.com/problems/reverse-nodes-in-k-group/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	left, right := reverse(head, k)
	if right == nil {
		return left
	}

	for right.Next != nil {
		next, nRight := reverse(right.Next, k)
		if nRight == nil {
			break
		}
		right.Next = next
		right = nRight
	}

	return left
}

func reverse(head *ListNode, k int) (*ListNode, *ListNode) {
	left, right := head, head

	for i := 1; i < k; i++ {
		if right.Next == nil {
			return left, nil
		}
		right = right.Next
	}

	lNext := left.Next
	rNext := right.Next

	for left != right {
		left.Next = rNext
		rNext = left
		left = lNext
		lNext = left.Next
	}

	left.Next = rNext
	return right, head
}
