package leetcode

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	for i, num := range nums {
		if dp[i] == 0 {
			continue
		}
		for j := i + 1; j <= i+num; j++ {
			if j >= len(nums) {
				continue
			}
			if dp[j] != 0 && dp[j] < dp[i]+1 {
				continue
			}
			dp[j] = dp[i] + 1
		}
	}
	return dp[len(dp)-1] - 1
}
