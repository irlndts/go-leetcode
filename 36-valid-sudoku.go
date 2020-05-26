package leetcode

// https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	// rows
	for _, b := range board {
		if !isValid(b) {
			return false
		}
	}

	// columns
	columns := make([]byte, 0, 9)
	for i := 0; i < 9; i++ {
		columns = []byte{}
		for j := 0; j < 9; j++ {
			columns = append(columns, board[j][i])
		}
		if !isValid(columns) {
			return false
		}
	}

	// squares
	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			block := make([]byte, 0, 9)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					block = append(block, board[i+k][j+l])
				}
			}
			if !isValid(block) {
				return false
			}
		}
	}

	return true
}

func isValid(nums []byte) bool {
	m := make(map[byte]struct{})
	for _, n := range nums {
		if _, ok := m[n]; ok && string(n) != "." {
			return false
		}
		m[n] = struct{}{}
	}

	return true
}
