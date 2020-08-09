package leetcode

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		a, b = equalize(a, b)
	} else if len(b) > len(a) {
		b, a = equalize(b, a)
	}

	result := ""
	var add bool
	fmt.Println(a, b)
	for i := len(a) - 1; i >= 0; i-- {
		switch {
		case add:
			switch {
			case string(a[i]) == "1" && string(b[i]) == "1":
				result = "1" + result
			case string(a[i]) == "1" || string(b[i]) == "1":
				result = "0" + result
			default:
				result = "1" + result
				add = false
			}
		case string(a[i]) == "1" && string(b[i]) == "1":
			add = true
			result = "0" + result
			fmt.Println(i, result)
		case string(a[i]) == "1" || string(b[i]) == "1":
			add = false
			result = "1" + result
		default:
			result = "0" + result
			add = false
		}
	}
	if add {
		result = "1" + result
	}

	return result
}

func equalize(a, b string) (string, string) {
	b = strings.Repeat("0", len(a)-len(b)) + b
	return a, b
}
