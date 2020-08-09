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

var numbers []string

func sumNumbers(root *TreeNode) int {
	numbers = []string{}
	result := 0
	sum(root, "")
	for _, n := range numbers {
		number, _ := strconv.Atoi(n)
		result += number
	}
	return result
}

func sum(root *TreeNode, number string) {
	if root == nil {
		return
	}

	number += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		numbers = append(numbers, number)
	}

	sum(root.Left, number)
	sum(root.Right, number)
}
