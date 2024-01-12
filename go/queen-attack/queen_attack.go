package queenattack

import "fmt"

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	err := testValid(whitePosition, blackPosition)
	if err != nil {
		return false, err
	}
	return true, nil
}

func testValid(whitePosition, blackPosition string) error {
	if blackPosition == "" || whitePosition == "" {
		return fmt.Errorf("position can't be empty")
	}
	if whitePosition[0] < 'a' || whitePosition[0] > 'h' || whitePosition[1] < '1' || whitePosition[1] > '8' {
		return fmt.Errorf("white position invalid")
	}
	if blackPosition[0] < 'a' || blackPosition[0] > 'h' || blackPosition[1] < '1' || blackPosition[1] > '8' {
		return fmt.Errorf("black position invalid")
	}
	if blackPosition == whitePosition {
		return fmt.Errorf("white and black position can't be the same")
    }
	return nil
}
