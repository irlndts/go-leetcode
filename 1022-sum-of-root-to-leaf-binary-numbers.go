package leetcode

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var sums []string

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sums = []string{}
	path(root, "")
	result := 0
	for _, sum := range sums {
		result += encoder(sum)
	}

	return result
}

func encoder(sum string) int {
	result := 0
	p := 1
	for i := len(sum) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(sum[i]))
		result += n * p
		p *= 2
	}

	return result
}

func path(root *TreeNode, sum string) {
	if root == nil {
		return
	}

	sum += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		sums = append(sums, sum)
	}

	path(root.Left, sum)
	path(root.Right, sum)
}
