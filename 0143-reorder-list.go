package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	list := []*ListNode{}
	start := head
	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	if len(list) <= 2 {
		head = start
		return
	}

	head = start
	head.Next = list[len(list)-1]
	head = head.Next

	for i := 1; i < len(list)/2; i++ {
		head.Next = list[i]
		head.Next.Next = list[len(list)-1-i]
		head = head.Next.Next
	}

	if len(list)%2 == 1 {
		head.Next = list[len(list)/2]
		head = head.Next
	}

	head.Next = nil
	head = start
}
