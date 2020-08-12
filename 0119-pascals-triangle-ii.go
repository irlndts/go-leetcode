package leetcode

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	prev := []int{1, 1}
	var cur []int
	for i := 1; i < rowIndex; i++ {
		cur = []int{}
		for j := 1; j < len(prev); j++ {
			cur = append(cur, prev[j-1]+prev[j])
		}
		cur = append([]int{1}, cur...)
		cur = append(cur, 1)
		prev = cur
	}

	return prev
}

func getRow1(rowIndex int) []int {
	line := rowIndex + 1
	val := 1
	res := make([]int, line)
	for i := 1; i <= line; i++ {
		res[i-1] = val
		val = val * (line - i) / i
	}
	return res
}
