package leetcode

func longestPalindrome(s string) int {
	result := 0
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	found := false
	for _, v := range m {
		if v%2 == 0 {
			result += v
		}
		if v%2 == 1 {
			result += v - 1
			found = true
		}
	}

	if found {
		result++
	}

	return result
}
