package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result = 0
	path(root)
	return result
}

func path(root *TreeNode) {
	if root == nil {
		return
	}

	cur := distance(root.Left, 0) + distance(root.Right, 0)
	result = max(result, cur)

	path(root.Left)
	path(root.Right)

}

func distance(root *TreeNode, current int) int {
	if root == nil {
		return current
	}
	current++
	return max(distance(root.Left, current), distance(root.Right, current))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var result = 0
