package leetcode

// https://leetcode.com/problems/maximum-depth-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	return visitNode(root)
}

func visitNode(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	l := 1 + visitNode(node.Left)
	r := 1 + visitNode(node.Right)
	if l > r {
		return l
	}
	return r
}
