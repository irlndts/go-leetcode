package problems

func plusOne(digits []int) []int {
	plus := false
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		if plus {
			digits[i]++
			plus = false
		}
		if digits[i] == 10 {
			digits[i] = 0
			plus = true
			continue
		}
		break
	}
	if plus {
		digits = append([]int{1}, digits...)
	}
	return digits
}
