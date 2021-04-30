package problems

import (
	"strings"
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)

	l, r := 0, len(s)-1
	for l <= r {
		if !unicode.IsDigit(rune(s[l])) && !unicode.IsLetter(rune(s[l])) {
			l++
			continue
		}

		if !unicode.IsDigit(rune(s[r])) && !unicode.IsLetter(rune(s[r])) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
