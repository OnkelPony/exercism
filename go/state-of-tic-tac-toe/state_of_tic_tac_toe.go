package stateoftictactoe

import "errors"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	switch {
	case iswin(board):
		return Win, nil
	case !isValid(board):
		return "", errors.New("Invalid board")
	case isFull(board):
		return Draw, nil
	default:
		return Ongoing, nil
	}
}
