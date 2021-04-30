package problems

func maxSubArray(nums []int) int {
	result := nums[0]
	for i := 0; i < len(nums); i++ {
		if got := max(nums[i:]); got > result {
			result = got
		}
	}
	return result
}

func max(nums []int) int {
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < sum+nums[i] {
			max = sum + nums[i]
		}
		sum += nums[i]
	}
	return max
}
