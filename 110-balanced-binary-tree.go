package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return balanced(root)
}

func balanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if diff := length(root.Left, 1) - length(root.Right, 1); diff > 1 || diff < -1 {
		return false
	}

	if !balanced(root.Left) {
		return false
	}
	if !balanced(root.Right) {
		return false
	}
	return true
}

func length(root *TreeNode, level int) int {
	if root == nil {
		return level
	}

	return max(length(root.Left, level+1), length(root.Right, level+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
