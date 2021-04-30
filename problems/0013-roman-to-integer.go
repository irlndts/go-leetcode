package problems

// https://leetcode.com/problems/roman-to-integer/

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0
	for i := range s {
		if i != len(s)-1 && m[string(s[i])] < m[string(s[i+1])] {
			result -= m[string(s[i])]
			continue
		}
		result += m[string(s[i])]
	}
	return result
}
