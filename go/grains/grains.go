// Package grains counts grains of wheat on chessboard.
package grains

import (
	"fmt"
)

// Square returns number of wheat grains on certain chess square.
func Square(squareNo int) (uint64, error) {
	if squareNo < 1 || squareNo > 64 {
		return 0, fmt.Errorf("bad square number %d, must be between 1 and 64", squareNo)
	}
	return 1 << (squareNo - 1), nil
}

// Total returns number of wheat grains on all chess squares together.
func Total() uint64 {
	return 1<<64 - 1
}
