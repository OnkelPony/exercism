package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// SillyNephewError exists for nephew errors of negative count of cows.
type SillyNephewError struct {
	noCows int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.noCows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	switch err {
	case nil:
		if cows == 0 {
			return 0, errors.New("Division by zero")
		} else if cows < 0 {
			return 0, &SillyNephewError{noCows: cows}
		} else if fodder < 0 {
			return 0, errors.New("Negative fodder")
		} else {
			return fodder / float64(cows), nil
		}
	case ErrScaleMalfunction:
		return 2 * fodder / float64(cows), nil
	default:
		return 0, err
	}
}
