package leetcode

// https://leetcode.com/problems/path-sum/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return pathSum(root, sum)
}

func pathSum(root *TreeNode, sum int) bool {
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			return true
		}
		return false
	}

	if root.Left != nil {
		if left := pathSum(root.Left, sum); left {
			return left
		}
	}

	if root.Right != nil {
		if right := pathSum(root.Right, sum); right {
			return right
		}
	}
	return false
}
