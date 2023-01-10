package problems

func islandPerimeter(grid [][]int) int {
	var result int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {

				if i == 0 || grid[i-1][j] == 0 {
					result++
				}

				if j == 0 || grid[i][j-1] == 0 {
					result++
				}

				if i == len(grid)-1 || grid[i+1][j] == 0 {
					result++
				}

				if j == len(grid[0])-1 || grid[i][j+1] == 0 {
					result++
				}
			}
		}
	}
	return result
}
