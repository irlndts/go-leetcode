package problems

// https://leetcode.com/problems/range-sum-query-2d-immutable/

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		matrix: matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := 0
	for i := row1; i <= row2; i++{
		for j := col1; j <= col2; j++{
			result += this.matrix[i][j]
		}
	}
	return result
}