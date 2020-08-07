package leetcode

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	diff := len(nums) / 2
	m := diff
	for diff != 0 {
		if target == nums[m] {
			return m
		}

		diff = diff / 2
		if target > nums[m] {
			m = m + diff
		} else {
			m = m - diff
		}
	}

	if m-1 >= 0 && target == nums[m-1] {
		return m - 1
	}

	if m-2 >= 0 && target == nums[m-2] {
		return m - 2
	}

	if m+1 < len(nums) && target > nums[m+1] {
		return m + 2
	}

	if m+2 < len(nums) && target > nums[m+2] {
		return m + 3
	}

	if target > nums[m] {
		return m + 1
	}

	return m
}
