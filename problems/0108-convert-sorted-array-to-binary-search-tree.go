package problems

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	median := len(nums) / 2
	node := &TreeNode{
		Val:   nums[median],
		Left:  sortedArrayToBST(nums[:median]),
		Right: sortedArrayToBST(nums[median+1:]),
	}

	return node
}
