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
			buf.WriteByte(board[row][col])
		}
		result[row] = buf.String()
	}
	return result
}
