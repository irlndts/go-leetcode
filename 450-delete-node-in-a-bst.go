package leetcode

// https://leetcode.com/problems/delete-node-in-a-bst/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	switch {
	case key < root.Val:
		root.Left = deleteNode(root.Left, key)
	case key > root.Val:
		root.Right = deleteNode(root.Right, key)
	default:
		return removeNode(root)
	}

	return root
}

func removeNode(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return nil
	}
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}

	node, previous := root.Right, root.Right
	for node.Left != nil {
		previous = node
		node = node.Left
	}

	previous.Left = nil
	if node.Right != nil {
		previous.Left = node.Right
	}
	node.Left = root.Left

	if node == root.Right {
		node.Right = root.Right.Right
		return node
	}

	node.Right = root.Right
	return node
}
