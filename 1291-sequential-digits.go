package leetcode

func sequentialDigits(low int, high int) []int {
	result := []int{}
	for _, d := range digits {
		if d >= low && d <= high {
			result = append(result, d)
		}
		if d > high {
			break
		}
	}
	return result
}

var digits []int

func init() {
	for i := 12; i < 90; i += 11 {
		digits = append(digits, i)
	}

	for i := 123; i < 800; i += 111 {
		digits = append(digits, i)
	}

	for i := 1234; i < 7000; i += 1111 {
		digits = append(digits, i)
	}

	for i := 12345; i < 60000; i += 11111 {
		digits = append(digits, i)
	}

	for i := 123456; i < 500000; i += 111111 {
		digits = append(digits, i)
	}

	for i := 1234567; i < 4000000; i += 1111111 {
		digits = append(digits, i)
	}

	for i := 12345678; i < 30000000; i += 11111111 {
		digits = append(digits, i)
	}

	for i := 123456789; i < 200000000; i += 111111111 {
		digits = append(digits, i)
	}
}
