package problems

// https://leetcode.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	arr := make([]string, numRows)
	line := 0
	up := true
	for i := 0; i < len(s); i++ {
		arr[line] += string(s[i])

		if line == (numRows - 1) {
			up = false
		} else if line == 0 {
			up = true
		}

		if up {
			line++
		} else {
			line--
		}

	}
	var result string
	for _, str := range arr {
		result += str
	}

	return result

}
