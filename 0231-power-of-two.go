package leetcode

// https://leetcode.com/problems/power-of-two/

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	for n != 1 {
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}
