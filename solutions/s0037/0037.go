package main

const ONES = 0x1ff

func SolveSudoku(board [][]byte) {
	rows := make([]int, 9)
	cols := make([]int, 9)
	grid := make([]int, 9)

	for i := range rows {
		rows[i], cols[i], grid[i] = ONES, ONES, ONES
	}

	for i, row := range board {
		for j, cell := range row {
			if cell < '1' || cell > '9' {
				continue
			}

			m := 1 << (cell - '1')
			k := i/3*3 + j/3

			rows[i] ^= m
			cols[j] ^= m
			grid[k] ^= m
		}
	}

	var solve func(n int) bool
	solve = func(n int) bool {
		if n == 81 {
			return true
		}

		i, j := n/9, n%9
		if '1' <= board[i][j] && board[i][j] <= '9' {
			return solve(n + 1)
		}

		for d := byte('1'); d <= '9'; d++ {
			m := 1 << (d - '1')
			k := i/3*3 + j/3

			if m&rows[i]&cols[j]&grid[k] == 0 {
				continue
			}

			board[i][j] = d
			rows[i] ^= m
			cols[j] ^= m
			grid[k] ^= m

			if solve(n + 1) {
				return true
			}

			board[i][j] = '.'
			rows[i] ^= m
			cols[j] ^= m
			grid[k] ^= m
		}

		return false
	}

	solve(0)
}
