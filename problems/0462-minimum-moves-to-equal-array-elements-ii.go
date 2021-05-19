package problems

// https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	mid := nums[len(nums)/2]

	result := 0
	for _, num := range nums{
		result += mod(mid - num)
	}
	return result
}

func mod(a int) int{
	if a < 0 {
		return -a
	}
	return a
}