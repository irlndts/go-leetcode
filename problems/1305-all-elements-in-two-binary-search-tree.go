package problems

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var elements []int

func getAllElements(root1, root2 *TreeNode) []int {
	elements = []int{}
	helper(root1)
	helper(root2)
	sort.Ints(elements)
	return elements
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	elements = append(elements, root.Val)
	helper(root.Left)
	helper(root.Right)
}
