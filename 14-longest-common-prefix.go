package leetcode

// https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	if len(result) == 0 {
		return ""
	}
	for i, str := range strs {
		if i == 0 {
			continue
		}
		if len(str) == 0 {
			return ""
		}
		result = compare(result, str)
		if result == "" {
			return result
		}
	}

	return result
}

func compare(origin, other string) string {
	for i, s := range origin {
		if len(other) == i {
			return origin[:i]
		}
		if s != rune(other[i]) {
			if i == 0 {
				return ""
			}
			return origin[:i]
		}
	}
	return origin
}
