package problems

func numTrees(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		for left := 0; left <= i-1; left++ {
			dp[i] += dp[left] * dp[i-1-left]
		}
	}
	return dp[n]
}
