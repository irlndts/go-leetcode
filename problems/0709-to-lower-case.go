package problems

// https://leetcode.com/problems/to-lower-case/

func toLowerCase(s string) string {
	result := ""
	for _, l := range s {
		if l >= 'A' && l <= 'Z' {
			result += string(l + 32)
		} else {
			result += string(l)
		}
	}

	return result
}
