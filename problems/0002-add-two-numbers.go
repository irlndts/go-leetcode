package problems

// https://leetcode.com/problems/add-two-numbers/

import "math/big"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := big.NewInt(0)
	sum.Add(getDigit(l1), getDigit(l2))

	if sum.Cmp(big.NewInt(0)) == 0 {
		return &ListNode{Val: 0}
	}

	// start point
	var v *big.Int
	div := big.NewInt(0)
	sum, v = sum.DivMod(sum, big.NewInt(10), div)
	result := &ListNode{Val: int(v.Int64())}
	entrance := result
	for sum.Cmp(big.NewInt(0)) != 0 {
		sum, v = sum.DivMod(sum, big.NewInt(10), div)
		result.Next = &ListNode{Val: int(v.Int64())}
		result = result.Next
	}
	return entrance
}

func getDigit(l *ListNode) *big.Int {
	result := big.NewInt(0)
	shift := big.NewInt(1)
	a := big.NewInt(0)
	for {
		a.Mul(big.NewInt(int64(l.Val)), shift)
		result.Add(result, a)
		shift.Mul(shift, big.NewInt(10))
		if l.Next == nil {
			return result
		}
		l = l.Next
	}
	return result
}
