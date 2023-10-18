package bottlesong

import (
	"fmt"
	"strings"
)

const (
	firstAndSecondRowOne = "green bottle hanging on the wall,"
	firstAndSecondRow    = "green bottles hanging on the wall,"
	thirdRow             = "And if one green bottle should accidentally fall,"
	fourthRowStart       = "There'll be"
	fourthRowEndOne      = "green bottle hanging on the wall."
	fourthRowEnd         = "green bottles hanging on the wall."
)

func Recite(startBottles, takeDown int) []string {
	numbers := []string{"no", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	var result []string
	for i := 0; i < takeDown; i++ {
		bottles := startBottles - i
		if bottles == 1 {
			result = append(result, fmt.Sprintf("%s %s", numbers[bottles], firstAndSecondRowOne))
			result = append(result, fmt.Sprintf("%s %s", numbers[bottles], firstAndSecondRowOne))
		} else {
			result = append(result, fmt.Sprintf("%s %s", numbers[bottles], firstAndSecondRow))
			result = append(result, fmt.Sprintf("%s %s", numbers[bottles], firstAndSecondRow))
		}
		result = append(result, thirdRow)
		if bottles == 2 {
			result = append(result, fmt.Sprintf("%s %s %s", fourthRowStart, strings.ToLower(numbers[bottles-1]), fourthRowEndOne))
		} else {
			result = append(result, fmt.Sprintf("%s %s %s", fourthRowStart, strings.ToLower(numbers[bottles-1]), fourthRowEnd))
		}
		if i+1 < takeDown {
			result = append(result, "")
		}
	}
	return result
}
