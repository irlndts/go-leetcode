package leetcode

func solveSudoku(board [][]byte) {
	path(board)
}

func path(board [][]byte) bool {
	var m, n int
	found := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				m, n = i, j
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return true
	}

	for _, i := range []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'} {
		if isValidNumber(board, m, n, i) {
			board[m][n] = i
			if path(board) {
				return true
			}
			board[m][n] = '.'
		}
	}

	return false
}

func isValidNumber(board [][]byte, i, j int, num byte) bool {
	// row
	for _, b := range board[i] {
		if b == num {
			return false
		}
	}

	// column
	for l := 0; l < 9; l++ {
		if board[l][j] == num {
			return false
		}
	}

	// square
	k := 6
	if i < 6 {
		k = 3
	}
	if i < 3 {
		k = 0
	}

	l := 9
	if j < 6 {
		l = 3
	}
	if j < 3 {
		l = 0
	}

	for m := k; m < k+3; m++ {
		for n := l; n < l+3; n++ {
			if board[m][n] == num {
				return false
			}
		}
	}

	return true
}
