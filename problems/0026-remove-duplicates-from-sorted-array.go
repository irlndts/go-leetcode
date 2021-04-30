package problems

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	l := len(nums)
	for i := 1; i < l; i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i-1], nums[i:]...)
			i--
			l--
		}
	}
	return len(nums)
}
