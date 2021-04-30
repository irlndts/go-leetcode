package problems

import "strconv"

func reverse(x int) int {
	str := strconv.Itoa(x)
	var negative bool
	if string(str[0]) == "-" {
		negative = true
		str = str[1:]
	}
	var result string
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	integer, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}
	if int(int32(integer)) != integer {
		return 0
	}
	if negative {
		return -integer
	}

	return integer
}
