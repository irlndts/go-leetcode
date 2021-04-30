package problems

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	result := make([]int, 0, len(matrix)*len(matrix[0]))

	layer := 0
	top, right, bottom, left := false, false, false, false
	for layer <= len(matrix)/2 {
		// top
		for i := layer; i < len(matrix[0])-layer; i++ {
			result = append(result, matrix[layer][i])
			top = true
		}

		if !top {
			break
		}

		// right
		for i := layer + 1; i < len(matrix)-layer; i++ {
			result = append(result, matrix[i][len(matrix[0])-layer-1])
			right = true
		}
		if !right {
			break
		}

		// bottom
		for i := len(matrix[0]) - layer - 2; i >= layer; i-- {
			result = append(result, matrix[len(matrix)-layer-1][i])
			bottom = true
		}
		if !bottom {
			break
		}

		// left
		for i := len(matrix) - layer - 2; i >= layer+1; i-- {
			result = append(result, matrix[i][layer])
			left = true
		}
		if !left {
			break
		}

		top, right, bottom, left = false, false, false, false
		layer++
	}

	return result
}
