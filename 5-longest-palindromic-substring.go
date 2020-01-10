package leetcode

// https://leetcode.com/problems/longest-palindromic-substring/

// the idea is to check substrings of the largest length of left and side parts
// and check them from the middle.
// for example, if the string is abeffg
// substrings are
// ab
// abe
// abef
// abeffg
// beffg
// effg - there is the longest palindrome in the middle for the example
// ffg
// fg

func longestPalindrome(l string) string {
	if len(l) < 2 {
		return l
	}

	result := string(l[0])
	left := 0
	right := 1
	for left != len(l)-1 {
		if right-left < len(result) {
			break
		}

		substr := l[left : right+1]
		if p := palindrome(substr); len(p) > len(result) {
			result = p
		}

		if right == len(l)-1 {
			left++
		} else {
			right++
		}
	}

	return result
}

func palindrome(s string) string {
	result := ""
	left := len(s)/2 - 1
	var right int
	if isOdd(len(s)) {
		right = len(s)/2 + 1
	} else {
		right = len(s) / 2
	}

	for {
		if left < 0 || right > len(s) {
			break
		}
		if s[left] != s[right] {
			break
		}

		result = s[left : right+1]
		left--
		right++
	}

	return result
}

func isOdd(l int) bool {
	if l%2 != 0 {
		return true
	}
	return false
}
