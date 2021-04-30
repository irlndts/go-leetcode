package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	results := []int{}
	lines = [][]int{}

	path(root, 0)

	for _, line := range lines {
		results = append(results, line[0])
	}

	return results
}

var lines [][]int

func path(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(lines) <= level {
		lines = append(lines, []int{root.Val})
	} else {
		lines[level] = append(lines[level], root.Val)
	}

	path(root.Right, level+1)
	path(root.Left, level+1)
}
