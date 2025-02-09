package main

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	p := 0
	r := len(grid[0])

	for i := 0; i < len(grid); i++ {
		for k := 0; k < r; k++ {
			if grid[i][k] != 1 {
				continue
			}

			p += 4

			if i > 0 && grid[i-1][k] == 1 {
				p -= 2
			}
			if k > 0 && grid[i][k-1] == 1 {
				p -= 2
			}

		}
	}

	return p
}
