package problems

func isMatch(s, p string) bool {
	dp := make([][]bool, len(p)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}
	dp[0][0] = true

	for i := 0; i < len(p); i++ {
		switch p[i] {
		case '*':
			if i == 0 {
				continue
			}
			for j := 0; j <= len(s); j++ {
				if !dp[i+1][j] {
					dp[i+1][j] = dp[i-1][j]
				}

				if j < len(s) && (p[i-1] == s[j] || p[i-1] == '.') && dp[i+1][j] {
					dp[i+1][j+1] = true
				}
			}

		default:
			for j := 0; j < len(s); j++ {
				switch p[i] {
				case s[j], '.':
					if !dp[i+1][j+1] {
						dp[i+1][j+1] = dp[i][j]
					}
				}
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
