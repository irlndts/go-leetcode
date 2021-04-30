package problems

import "strconv"

// https://leetcode.com/problems/binary-watch/

func readBinaryWatch(num int) []string {
	vars = []string{}
	foo(num, 0, "")

	return bar(vars)
}

func bar(input []string) []string {
	result := make([]string, 0, len(input))
	for _, s := range input {
		hour := 0
		multiplier := 1
		for i := 3; i >= 0; i-- {
			if s[i] == '1' {
				hour += multiplier
			}
			multiplier *= 2
		}
		if hour > 11 {
			continue
		}

		minutes := 0
		multiplier = 1
		for i := len(s) - 1; i >= 4; i-- {
			if s[i] == '1' {
				minutes += multiplier
			}
			multiplier *= 2
		}
		if minutes > 59 {
			continue
		}

		current := strconv.Itoa(hour) + ":"
		if minutes < 10 {
			current += "0" + strconv.Itoa(minutes)
		} else {
			current += strconv.Itoa(minutes)
		}
		result = append(result, current)
	}
	return result
}

var vars []string

func foo(num, cur int, line string) {
	if len(line) == 10 {
		if num == cur {
			vars = append(vars, line)
		}
		return
	}

	foo(num, cur, line+"0")
	if cur != num {
		foo(num, cur+1, line+"1")
	}
}
