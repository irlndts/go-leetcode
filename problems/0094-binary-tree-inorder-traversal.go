package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack []*TreeNode

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	if s.IsEmpty() {
		return nil
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

func inorderTraversal(root *TreeNode) []int {

}
