package leetcode

func rotate(matrix [][]int) {
	level := 0

	for level < len(matrix)/2 {
		for i := level; i < len(matrix)-level-1; i++ {
			matrix[level][i],
				matrix[i][len(matrix)-level-1],
				matrix[len(matrix)-level-1][len(matrix)-i-1],
				matrix[len(matrix)-i-1][level] =
				matrix[len(matrix)-i-1][level],
				matrix[level][i],
				matrix[i][len(matrix)-level-1],
				matrix[len(matrix)-level-1][len(matrix)-i-1]
		}
		level++
	}
}
