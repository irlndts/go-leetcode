package problems

// https://leetcode.com/problems/swap-nodes-in-pairs/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	i := 0
	start := head.Next
	var previous *ListNode

	for head.Next != nil {

		if i%2 == 0 {
			current, next := head, head.Next.Next
			head = head.Next
			head.Next = current
			head.Next.Next = next
			if previous != nil {
				previous.Next = head
			}
		}

		i++
		previous = head
		head = head.Next
	}
	return start
}
