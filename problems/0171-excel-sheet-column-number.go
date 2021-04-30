package problems

func titleToNumber(s string) int {
	result := 0
	mul := 1
	for i := 0; i < len(s)-1; i++ {
		mul *= 26
	}
	for _, r := range s {
		v := r - 'A' + 1
		result += int(v) * mul
		mul /= 26
	}
	return result
}
