package leetcode

// https://leetcode.com/problems/4sum/

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	var result [][]int
	sort.Ints(nums)

	unique := map[string]struct{}{}
	positive := target >= 0

	for i := 0; i < len(nums)-3; i++ {
		one := nums[i]
		if positive && one > target {
			break
		}
		for j := i + 1; j < len(nums)-2; j++ {
			two := one + nums[j]
			if positive && two > target {
				break
			}
			for k := j + 1; k < len(nums)-1; k++ {
				three := two + nums[k]
				if positive && three > target {
					break
				}
				for l := k + 1; l < len(nums); l++ {
					sum := three + nums[l]
					if sum == target {
						s := fmt.Sprintf("%d:%d:%d:%d", nums[i], nums[j], nums[k], nums[l])
						if _, ok := unique[s]; ok {
							continue
						}
						unique[s] = struct{}{}

						result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
						continue
					}
					if positive && sum > target {
						break
					}
				}
			}
		}
	}

	return result
}
