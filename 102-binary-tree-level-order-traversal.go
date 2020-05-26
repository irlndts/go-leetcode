package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var results [][]int

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	results = [][]int{}

	level(root, 0)
	return results
}

func level(root *TreeNode, index int) {
	if root == nil {
		return
	}
	if len(results) < index+1 {
		results = append(results, []int{})
	}
	results[index] = append(results[index], root.Val)
	level(root.Left, index+1)
	level(root.Right, index+1)
}
