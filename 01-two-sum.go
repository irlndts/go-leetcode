package leetcode

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	checked := make(map[int]bool)
	for i := 0; i < len(nums)-1; i++ {
		if _, ok := checked[i]; ok {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
		checked[i] = true
	}
	return nil
}
