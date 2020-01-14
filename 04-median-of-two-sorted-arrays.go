package leetcode

import (
	"sort"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	return median(nums)
}

func median(nums []int) float64 {
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2
	}
	return float64(nums[len(nums)/2])
}
