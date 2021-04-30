package problems

// https://leetcode.com/problems/number-of-rectangles-that-can-form-the-largest-square/

func countGoodRectangles(rectangles [][]int) int {
	result := 0
	max := 0
	for _, rect := range rectangles {
		m := min(rect[0], rect[1])
		if m > max {
			result = 1
			max = m
		} else if m == max {
			result++
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
