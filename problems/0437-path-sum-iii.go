package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	result = 0
	summ(root, sum)
	return result
}

func summ(root *TreeNode, target int) {
	if root == nil {
		return
	}

	path(root, target, 0)

	summ(root.Left, target)
	summ(root.Right, target)
}

var result int

func path(root *TreeNode, target, current int) {
	if root == nil {
		return
	}

	current += root.Val
	if current == target {
		result++
	}

	path(root.Left, target, current)
	path(root.Right, target, current)

}
