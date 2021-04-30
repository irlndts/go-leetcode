package problems

// https://leetcode.com/problems/merge-two-sorted-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var start *ListNode
	if l1.Val <= l2.Val {
		start = &ListNode{
			Val: l1.Val,
		}
		l1 = l1.Next
	} else {
		start = &ListNode{
			Val: l2.Val,
		}
		l2 = l2.Next
	}

	entrance := start
	for {
		if l1 == nil || l2 == nil {
			break
		}
		if l1.Val <= l2.Val {
			start.Next = l1
			l1 = l1.Next
		} else {
			start.Next = l2
			l2 = l2.Next
		}
		start = start.Next
	}
	if l1 != nil {
		start.Next = l1
	}
	if l2 != nil {
		start.Next = l2
	}

	return entrance
}
