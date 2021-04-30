package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
	result = 0
	up(root)

	return result
}

var result int

func up(root *TreeNode) {
	if root == nil {
		return
	}
	ll, sl, isLeft := sum(root.Left)
	lr, sr, isRight := sum(root.Right)

	if ll == sl && lr == sr && isLeft && isRight {
		return
	}

	if ll < sl {
		root.Val += sl - ll
		root.Left.Val -= sl - ll
		result += sl - ll
	}
	if lr < sr {
		root.Val += sr - lr
		root.Right.Val -= sr - lr
		result += sr - lr
	}

	if ll > sl {
		root.Left.Val += ll - sl
		root.Val -= ll - sl
		result += ll - sl
	}

	if lr > sr {
		root.Right.Val += lr - sr
		root.Val -= lr - sr
		result += lr - sr
	}

	up(root.Left)
	up(root.Right)
}

func sum(root *TreeNode) (int, int, bool) {
	if root == nil {
		return 0, 0, true
	}

	l1, l2, isOne1 := sum(root.Left)
	r1, r2, isOne2 := sum(root.Right)
	return 1 + l1 + r1, root.Val + l2 + r2, root.Val == 1 && isOne1 && isOne2
}
