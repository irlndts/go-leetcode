package leetcode

// https://leetcode.com/problems/power-of-three/

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n % 3 == 0 {
		return isPowerOfThree(n/3)
	}
	return n==1
}
