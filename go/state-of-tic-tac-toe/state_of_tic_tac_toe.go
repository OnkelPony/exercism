package stateoftictactoe

import (
	"errors"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	switch {
	case !isValid(board):
		return "", errors.New("Invalid board")
	case iswin(board):
		return Win, nil
	case isFull(board):
		return Draw, nil
	default:
		return Ongoing, nil
	}
}

func isValid(board []string) bool {
	if countElement(board, 'O') > countElement(board, 'X') ||
		countElement(board, 'O')+1 < countElement(board, 'X')||
		threeSymbols(board, 'X') && threeSymbols(board, 'O'){
		return false
	}
	return true
}

func countElement(board []string, r rune) int {
	var count int
	for _, row := range board {
		for _, element := range row {
			if element == r {
				count++
			}
		}
	}
	return count
}

func iswin(board []string) bool {
	return threeSymbols(board, 'X') != threeSymbols(board, 'O')
}

func threeSymbols(board []string, r byte) bool {
	var inFalling, inRising int
	for i := range board {
		var inRow, inCol int
		for j := range board[i] {
			if board[i][j] == r {
				inRow++
				if i == j {
					inFalling++
				}
				}
			if board[j][i] == r {
				inCol++
				if i + j == len(board)-1 {
					inRising++
				}
			}
		}
		if inRow == len(board) || inCol == len(board) {
			return true
		}
	}
	return inRising == len(board) || inFalling == len(board)
}

func isFull(board []string) bool {
	return countElement(board, 'X') + countElement(board, 'O') == len(board) * len(board[0])
}
