package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
	subtreeWeight int
	result        int
)

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result = root.Val
	nodeWeight(root)
	return result
}

func nodeWeight(root *TreeNode) {
	if root == nil {
		return
	}

	weight := root.Val
	if root.Left != nil {
		subtreeWeight = root.Left.Val
		subtreeMaxWeight(root.Left, 0)
		if subtreeWeight > 0 {
			weight += subtreeWeight
		}
	}

	if root.Right != nil {
		subtreeWeight = root.Right.Val
		subtreeMaxWeight(root.Right, 0)
		if subtreeWeight > 0 {
			weight += subtreeWeight
		}
	}

	if weight > result {
		result = weight
	}

	nodeWeight(root.Left)
	nodeWeight(root.Right)
}

func subtreeMaxWeight(root *TreeNode, weight int) {
	if root == nil {
		return
	}

	weight += root.Val
	if weight > subtreeWeight {
		subtreeWeight = weight
	}

	subtreeMaxWeight(root.Left, weight)
	subtreeMaxWeight(root.Right, weight)
}
