// Package allyourbase implements conversion of numbers between various number bases.
package allyourbase

import (
	"errors"
	"math"
)

// ConvertToBase returns number in outputBase, equal to number in input base in argument.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	if len(inputDigits) == 0 {
		return []int{0}, nil
	}
	maxIndex := len(inputDigits) - 1
	var decaResult int
	for i, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		decaResult += digit * int(math.Pow(float64(inputBase), float64(maxIndex-i)))
	}
	if decaResult == 0 {
		return []int{0}, nil
	}
	var outDigitsLength int
	for j := 0; ; j++ {
		if math.Pow(float64(outputBase), float64(j)) > float64(decaResult) {
			outDigitsLength = j
			break
		}
	}
	result := make([]int, outDigitsLength)
	for k := 0; k < outDigitsLength; k++ {
		divider := int(math.Pow(float64(outputBase), float64(outDigitsLength-1-k)))
		result[k] = decaResult / divider
		decaResult %= divider
	}
	return result, nil
}
