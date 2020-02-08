package leetcode

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

	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}

	median := len(nums) / 2

	node := &TreeNode{
		Val:  nums[median],
		Left: sortedArrayToBST(nums[:median]),
	}
	if len(nums)-1 != median {
		node.Right = sortedArrayToBST(nums[median+1:])
	}

	return node
}
