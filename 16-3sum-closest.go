package leetcode

// https://leetcode.com/problems/3sum-closest/submissions/

import "sort"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	dif := mod(result - target)

	l, r := 0, 0

	for i := 0; i < len(nums)-2; i++ {
		l, r = i+1, len(nums)-1
		curDif := mod((nums[i] + nums[l] + nums[r]) - target)
		for l < r {
			v := nums[i] + nums[l] + nums[r]
			if v == target {
				return v
			}

			if newCurDif := mod(v - target); newCurDif < curDif {
				curDif = newCurDif
			}

			if curDif < dif {
				dif = curDif
				result = v
			}

			if v < target {
				l++
				continue
			}

			r--
		}
	}

	return result
}

func mod(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
