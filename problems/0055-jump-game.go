package problems

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		if !dp[i] {
			continue
		}

		for j := i; j <= i+nums[i]; j++ {
			if j >= len(nums) {
				break
			}
			dp[j] = true
		}
	}
	return dp[len(nums)-1]
}
