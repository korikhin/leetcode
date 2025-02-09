package main

func gameOfLife(board [][]int) {
	ROWS, COLS := len(board), len(board[0])
	countNeighbors := func(x, y int) int {
		count := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}

				row := x + i
				col := y + j
				if row < 0 || row >= ROWS || col < 0 || col >= COLS {
					continue
				}

				count += board[row][col]
			}
		}
		return count
	}

	newBoard := make([][]int, ROWS)
	for i := range newBoard {
		newBoard[i] = make([]int, COLS)
	}

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			neighbors := countNeighbors(i, j)
			if board[i][j] == 1 {
				if neighbors == 2 || neighbors == 3 {
					newBoard[i][j] = 1
				}
			} else {
				if neighbors == 3 {
					newBoard[i][j] = 1
				}
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = newBoard[i][j]
		}
	}
}
