package problems

import "strings"

// https://leetcode.com/problems/implement-strstr/

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
