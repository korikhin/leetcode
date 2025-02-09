package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var (
		rows = len(grid) - 1
		cols = len(grid[0]) - 1

		islands int
		dfs     func(row, col int)
	)

	dfs = func(row, col int) {
		if row < 0 || col < 0 || row > rows || col > cols {
			return
		}
		if grid[row][col] != '1' {
			return
		}

		// Mark as visited
		grid[row][col] = '0'

		dfs(row, col-1)
		dfs(row, col+1)
		dfs(row-1, col)
		dfs(row+1, col)
	}

	// Make sure to traverse through each cell
	for i := range grid {
		for k := range grid[0] {
			if grid[i][k] == '1' {
				islands++
				dfs(i, k)
			}
		}
	}

	return islands
}
