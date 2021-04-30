package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	memo = map[*TreeNode]int{}

	result = 0
	path(root)
	return result
}

var memo map[*TreeNode]int

func path(root *TreeNode) {
	if root == nil {
		return
	}

	cur := distance(root.Left) + distance(root.Right)
	result = max(result, cur)

	path(root.Left)
	path(root.Right)

}

func distance(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if v, ok := memo[root]; ok {
		return v
	}

	memo[root] = 1 + max(distance(root.Left), distance(root.Right))
	return memo[root]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var result = 0
