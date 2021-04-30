package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	left, right := headA, headB
	for left != right {
		if left == nil {
			left = headB
		} else {
			left = left.Next
		}

		if right == nil {
			right = headA
		} else {
			right = right.Next
		}
	}
	return left
}
