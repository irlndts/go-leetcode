package leetcode

import (
	"math"
	"strconv"
)

// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(str string) int {
	var minus, digit bool

	result := ""
	for _, s := range str {
		if s == ' ' && !digit {
			continue
		}
		if s == '-' && !digit {
			digit = true
			minus = true
			continue
		}
		if s == '+' && !digit {
			digit = true
			continue
		}

		if s < '0' || s > '9' {
			break
		}

		digit = true
		result += string(s)
	}

	res, _ := strconv.Atoi(result)
	if minus {
		res = -res
	}
	if res < math.MinInt32 {
		return math.MinInt32
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}
