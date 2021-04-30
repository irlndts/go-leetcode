package problems

var output int
var found bool

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	output = 0
	found = false
	visited := make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[0]))
	}
	path(grid, visited, 0, 0, 0)

	return output
}

func path(grid, visited [][]int, i, j int, result int) {
	if visited[i][j] != 0 && visited[i][j] <= result {
		return
	}

	if len(grid)-1 == i && len(grid[0])-1 == j {
		if !found {
			output = result + grid[i][j]
			found = true
			return
		}

		if result+grid[i][j] < output {
			output = result + grid[i][j]
			found = true
		}
		return
	}

	visited[i][j] = result + grid[i][j]

	if j < len(grid[0])-1 {
		path(grid, visited, i, j+1, result+grid[i][j])
	}

	if i < len(grid)-1 {
		path(grid, visited, i+1, j, result+grid[i][j])
	}

}
