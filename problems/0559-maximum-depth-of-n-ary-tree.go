package problems

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	result = 0
	path(root, 1)
	return result
}

var result int

func path(root *Node, level int) {
	if len(root.Children) == 0 {
		if level > result {
			result = level
		}
		return
	}

	for _, children := range root.Children {
		path(children, level+1)
	}
}
