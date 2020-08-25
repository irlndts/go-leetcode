package leetcode

func partitionDisjoint(A []int) int {
	dpMax := make([]int, len(A))
	dpMax[0] = A[0]
	for i := 1; i < len(A); i++ {
		dpMax[i] = max(A[i], dpMax[i-1])
	}

	dpMin := make([]int, len(A))
	dpMin[len(dpMin)-1] = A[len(A)-1]
	for i := len(A) - 2; i >= 0; i-- {
		dpMin[i] = min(A[i], dpMin[i+1])
	}

	for i := 1; i < len(A); i++ {
		if dpMax[i-1] <= dpMin[i] {
			return i
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
