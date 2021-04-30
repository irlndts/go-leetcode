package problems

// https://leetcode.com/problems/merge-k-sorted-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) != 1 {
		lists = mergePairs(lists)
	}

	return lists[0]
}

func mergePairs(lists []*ListNode) []*ListNode {
	if len(lists) < 2 {
		return lists
	}

	n := len(lists)
	if len(lists)%2 != 0 {
		n--
	}

	j := 0
	for i := 0; i < n; i += 2 {
		lists[j] = mergeLists(lists[i], lists[i+1])
		j++
	}

	if len(lists) != n {
		lists[j] = lists[len(lists)-1]
		j++
	}

	return lists[:j]
}

func mergeLists(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	var start, current *ListNode
	if a.Val <= b.Val {
		start = a
		a = a.Next
	} else {
		start = b
		b = b.Next
	}

	current = start
	for a != nil && b != nil {
		if a.Val < b.Val {
			current.Next = a
			current = current.Next
			a = a.Next
			continue
		}

		current.Next = b
		current = current.Next
		b = b.Next
	}

	if a != nil {
		current.Next = a
	}

	if b != nil {
		current.Next = b
	}

	return start
}
