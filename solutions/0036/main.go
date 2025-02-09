package main

func isValidSudoku(board [][]byte) bool {
	rows := make([]int, 9)
	cols := make([]int, 9)
	grid := make([]int, 9)

	for i, row := range board {
		for j, cell := range row {
			if cell < '1' || cell > '9' {
				continue
			}

			m := 1 << (cell - '1')
			k := i/3*3 + j/3

			if m&(rows[i]|cols[j]|grid[k]) != 0 {
				return false
			}

			rows[i] ^= m
			cols[j] ^= m
			grid[k] ^= m
		}
	}

	return true
}
