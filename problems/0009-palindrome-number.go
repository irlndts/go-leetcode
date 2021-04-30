package problems

import "strconv"

// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	rev := make([]string, 0, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		rev = append(rev, string(str[i]))
	}
	for i, s := range str {
		if string(s) != rev[i] {
			return false
		}
	}
	return true
}
