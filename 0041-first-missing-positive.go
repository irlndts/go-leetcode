package leetcode

import "sort"

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}
		if start == nums[i] {
			continue
		}
		if start+1 < nums[i] {
			return start + 1
		}
		start++
	}
	return start + 1
}
