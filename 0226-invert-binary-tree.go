package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	return revert(root)
}

func revert(root *TreeNode) *TreeNode {
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		root.Left = revert(root.Left)
	}
	if root.Right != nil {
		root.Right = revert(root.Right)
	}
	return root
}
