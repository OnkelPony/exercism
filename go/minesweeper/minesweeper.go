package minesweeper

import (
	"bytes"
)

// Annotate returns an annotated board
func Annotate(board []string) []string {
	result := make([]string, len(board))
	for row := range board {
		var buf bytes.Buffer
		for col := range board[row] {
			if board[row][col] == ' ' {
				buf.WriteByte(countMines(board, row, col))
			} else {
				buf.WriteByte(board[row][col])
			}
		}
		result[row] = buf.String()
	}
	return result
}

func countMines(board []string, row, col int) byte{
	var result int
	directions := [][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	h := len(board)
	w := len(board[0])
	for _, direction := range directions {
		position := [2]int{row + direction[0], col + direction[1]}
		if position[0] >= 0 && position[0] < h && position[1] >= 0 && position[1] < w {
			if board[position[0]][position[1]] == byte('*') {
				result++
			}
		}
	}
	if result == 0 {
		return byte(' ')
	}
	return byte(result + 48)
}
