package problems

import "strings"

func lengthOfLastWord(s string) int {
	line := strings.TrimRight(s, " ")
	i, l := 0, 0
	result := 0
	for _, letter := range line {
		if string(letter) == " " {
			l = i + 1
		}
		i++
	}

	result = i - l
	return result
}
