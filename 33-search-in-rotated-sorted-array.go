package leetcode

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		if nums[l] == target {
			return l
		}
		if nums[r] == target {
			return r
		}

		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[l] > nums[r] {
			l++
			r--
			continue
		}

		if nums[mid] < target {
			l = mid + 1
			continue
		}

		r = mid - 1
	}
	return -1
}
