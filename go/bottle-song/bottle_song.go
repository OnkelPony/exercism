package bottlesong

import (
	"fmt"
	"strings"
)

const (
	firstAndSecondRow = "green bottles hanging on the wall,"
	thirdRow = "And if one green bottle should accidentally fall,"
	fourthRowStart = "There'll be"
	fourthRowEndOne = "green bottle hanging on the wall."
	fourthRowEnd = "green bottles hanging on the wall."
)

func Recite(startBottles, takeDown int) []string {
	numbers := []string{"no", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	var result []string
	for i := 0; i < takeDown; i++ {
		result = append(result, fmt.Sprintf("%s %s", numbers[startBottles - i], firstAndSecondRow))
		result = append(result, fmt.Sprintf("%s %s", numbers[startBottles - i], firstAndSecondRow))
		result = append(result, thirdRow)
		result = append(result, fmt.Sprintf("%s %s %s", fourthRowStart, strings.ToLower(numbers[startBottles - i - 1]), fourthRowEnd))
		if i + 1 < takeDown {
			result = append(result, "")
		}
	}
	return result
}
