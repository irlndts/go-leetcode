package problems

import (
	"strconv"
	"strings"
)

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		tmp := strconv.Itoa(n % 1000)
		for len(tmp) != 3 {
			tmp = "0" + tmp
		}
		result = "." + tmp + result
		n /= 1000
	}
	result = strings.TrimLeft(result, ".0")
	return result
}
