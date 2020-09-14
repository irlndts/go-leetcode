package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0], dp[1], dp[2] = nums[0], nums[1], nums[2]+nums[0]

	for i := 3; i < len(nums); i++ {
		dp[i] = nums[i] + max(dp[i-3], dp[i-2])
	}

	return max(dp[len(dp)-1], dp[len(dp)-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
