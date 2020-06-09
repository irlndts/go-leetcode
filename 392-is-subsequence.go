package leetcode

// https://leetcode.com/problems/is-subsequence/

func isSubsequence(s string, t string) bool {
	index := 0
	found := false
	for _, r := range s {
		for j := index; j < len(t); j++ {
			if r == rune(t[j]) {
				index = j + 1
				found = true
				break
			}
		}
		if !found {
			return false
		}
		found = false
	}
	return true
}
