package minesweeper

// Annotate returns an annotated board
func Annotate(board []string) []string {
	result := make([]string, len(board))
	for r, row := range board {
		rowOut := []byte(row)
		for c := range board[r] {
			if board[r][c] == ' ' {
				rowOut[c] = (countMines(board, r, c))
			}
		}
		result[r] = string(rowOut)
	}
	return result
}

func countMines(board []string, row, col int) byte {
	var result byte
	directions := [][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	for _, direction := range directions {
		position := [2]int{row + direction[0], col + direction[1]}
		if position[0] >= 0 && position[0] < len(board) && position[1] >= 0 && position[1] < len(board[0]) {
			if board[position[0]][position[1]] == '*' {
				result++
			}
		}
	}
	if result == 0 {
		return ' '
	}
	return result + '0'
}
