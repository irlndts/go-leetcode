package leetcode

func binaryGap(N int) int {
	result, last := 0, -1

	for i := 0; i < 32; i++ {
		if (N>>i)&1 > 0 {
			if last >= 0 {
				result = max(result, i-last)
			}
			last = i
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
