package problems

// https://leetcode.com/problems/pascals-triangle/

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	var result [][]int
	result = append(result, []int{1})

	for i := 2; i <= numRows; i++ {
		for j := 0; j < i; j++ {
			if j == 0 {
				result = append(result, []int{1})
				continue
			}
			if j == i-1 {
				result[i-1] = append(result[i-1], 1)
				continue
			}
			result[i-1] = append(result[i-1], result[i-2][j-1]+result[i-2][j])
		}
	}

	return result
}
