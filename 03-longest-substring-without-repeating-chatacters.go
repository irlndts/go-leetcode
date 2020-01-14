package leetcode

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if result > len(s[i:]) {
			return result
		}
		m := make(map[byte]bool)
		row := 0
		for j := i; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				break
			}
			m[s[j]] = true
			row++
		}
		if row > result {
			result = row
		}
	}
	return result
}
