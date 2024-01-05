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
	if result == 0 {
		return byte(' ')
	}
	return byte(result + 48)
}
