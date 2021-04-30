package problems

// https://leetcode.com/problems/find-the-highest-altitude/

func largestAltitude(gain []int) int {
	current, max := 0, 0
	for _, g := range gain {
		current += g
		if current > max {
			max = current
		}
	}

	return max
}
