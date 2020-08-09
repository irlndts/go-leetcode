package leetcode

// https://leetcode.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == val {
			nums = pop(nums, i)
			l--
			i--
		}
	}
	return len(nums)
}

func pop(nums []int, i int) []int {
	if i == len(nums)-1 {
		return nums[:i]
	}
	return append(nums[:i], nums[i+1:]...)
}
