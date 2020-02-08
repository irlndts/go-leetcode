package leetcode

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	start, left, remove := head, head, head

	i := 0

	for {
		if i >= n {
			left = remove
			remove = remove.Next
		}

		if head.Next == nil {
			break
		}
		head = head.Next
		i++
	}

	// special case
	if i+1 == n && left == remove {
		return remove.Next
	}

	left.Next = remove.Next

	return start
}
