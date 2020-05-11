package leetcode

import "math"

// https://leetcode.com/problems/divide-two-integers/

func divide(dividend int, divisor int) int {
	r := math.Exp(math.Log(math.Abs(float64(dividend))) - math.Log(math.Abs(float64(divisor))))

	result := int(r)
	if (dividend < 0 || divisor < 0) && !(dividend < 0 && divisor < 0) {
		result = -result
	}

	if result < math.MinInt32 {
		return math.MinInt32
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	}

	return result
}
