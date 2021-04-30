package problems

func generateMatrix(n int) [][]int {
	layer := 0
	number := 1
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	for layer <= n/2 {
		for i := layer; i < n-layer; i++ {
			if result[layer][i] != 0 {
				break
			}
			result[layer][i] = number
			number++
		}

		for i := layer + 1; i < n-layer; i++ {
			if result[i][n-layer-1] != 0 {
				break
			}
			result[i][n-layer-1] = number
			number++
		}

		for i := n - layer - 2; i >= layer; i-- {
			if result[n-layer-1][i] != 0 {
				break
			}
			result[n-layer-1][i] = number
			number++
		}

		for i := n - layer - 2; i >= layer+1; i-- {
			if result[i][layer] != 0 {
				break
			}
			result[i][layer] = number
			number++
		}

		layer++
	}
	return result
}
