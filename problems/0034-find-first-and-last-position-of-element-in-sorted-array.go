package problems

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	l := search(nums, target, true)
	r := search(nums, target, false)
	return []int{l, r}
}

func search(nums []int, target int, left bool) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1

	for l <= r {
		m := (r + l) / 2
		if target == nums[m] {
			if left && (m == 0 || nums[m-1] != target) {
				return m
			}

			if !left && (m == len(nums)-1 || nums[m+1] != target) {
				return m
			}
		}

		if target > nums[m] || (target == nums[m] && !left) {
			l = m + 1
			continue
		}

		r = m - 1
	}

	return -1
}
