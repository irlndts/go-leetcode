package leetcode

import "sort"

// https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		need, left, right := -nums[i], i+1, len(nums)-1

		for left < right {
			sum := nums[left] + nums[right]
			if need == sum {
				result = append(result, []int{-need, nums[left], nums[right]})
				left++
				right--
				for left < right {
					if nums[left] != nums[left-1] {
						break
					}
					left++
				}

				for right > left {
					if nums[right] != nums[right+1] {
						break
					}
					right--
				}
				continue
			}

			if need < sum {
				right--
				continue
			}

			if need > sum {
				left++
				continue
			}

		}

	}

	return result
}
