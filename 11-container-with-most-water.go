package leetcode

// https://leetcode.com/problems/container-with-most-water/
// O(n)

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1

	result := 0
	sum, min := 0, 0

	for i != j {
		if height[i] < height[j] {
			min = height[i]
		} else {
			min = height[j]
		}

		sum = (j - i) * min
		if sum > result {
			result = sum
		}

		if height[i] > height[j] {
			j--
			continue
		}
		i++
	}

	return result
}
