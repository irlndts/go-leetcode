package leetcode

// https://leetcode.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	for _, r := range t {
		if _, ok := m[r]; !ok {
			return false
		}
		m[r]--
		if m[r] == 0 {
			delete(m, r)
		}
	}

	return true
}
