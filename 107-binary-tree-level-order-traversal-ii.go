package leetcode

// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var result [][]int

func levelOrderBottom(root *TreeNode) [][]int {
	result = [][]int{}
	if root == nil {
		return result
	}
	orderBottom(root, 0)

	// reverse
	r := make([][]int, len(result))
	for i := 0; i < len(result); i++ {
		r[i] = result[len(result)-i-1]
	}

	return r
}

func orderBottom(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(result) < level+1 {
		result = append(result, []int{})
	}

	result[level] = append(result[level], root.Val)
	orderBottom(root.Left, level+1)
	orderBottom(root.Right, level+1)
}
