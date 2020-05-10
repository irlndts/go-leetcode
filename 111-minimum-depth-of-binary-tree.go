package leetcode

// https://leetcode.com/problems/minimum-depth-of-binary-tree/

import "sort"

var result []int

func minDepth(root *TreeNode) int {
	result = []int{}

	if root == nil {
		return 0
	}

	depth(root, 1)

	sort.Ints(result)
	return result[0]
}

func depth(root *TreeNode, level int) {
	if root.Left == nil && root.Right == nil {
		result = append(result, level)
		return
	}

	if root.Left != nil {
		depth(root.Left, level+1)
	}
	if root.Right != nil {
		depth(root.Right, level+1)
	}
}
