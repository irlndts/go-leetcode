package leetcode

func maxTurbulenceSize(A []int) int {
	if len(A) <= 1 {
		return len(A)
	}

	if len(A) == 2 {
		if A[0] == A[1] {
			return 1
		}
		return 2
	}

	result := 1

	dp := make([]int, len(A))

	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if A[i] < A[i-1] {
			dp[i] = dp[i-1] - 1
		}
		if A[i] == A[i-1] {
			dp[i] = dp[i-1]
		}
	}

	for i := 0; i < len(dp)-1; i++ {
		l, r := dp[i], dp[i+1]
		if l == r {
			continue
		}

		index := 0
		cur := 0
		for j := i; j < len(dp); j++ {
			if (index%2 == 0 && dp[j] == l) ||
				(index%2 == 1 && dp[j] == r) {
				cur++
				index++
				if cur > result {
					result = cur
				}
				continue
			}

			if cur > result {
				result = cur
			}
			break
		}
	}

	return result
}
