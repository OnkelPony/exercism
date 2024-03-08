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
		countElement(board, 'O')+1 < countElement(board, 'X') {
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
	return true
}

func isFull(board []string) bool {
	return true
}
