// Package armstrong checks if provided number is Armstrong number.
package armstrong

import (
	"math"
	"strconv"
)

// IsNumber returns true, if argument is Armstrong number.
func IsNumber(input int) bool {
	orig := input
	power := len(strconv.Itoa(input))
	if power == 1 {
		return true
	}
	if power == 2 {
		return false
	}
	var sum int
	for divider := int(math.Pow(10, float64(power-1))); divider >= 1; divider /= 10 {
		sum += int(math.Pow(float64(input/divider), float64(power)))
		input %= divider
	}
	return sum == orig
}
