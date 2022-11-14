// Package hamming provides method to calculate the Hamming Distance between two DNA strands.
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance calculates the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("input strings must have equal length")
	}
	return calculate(a, b), nil
}

// Calculate takes two strings and returns count of their differences.
func calculate(a string, b string) int {
	aRunes := []rune(a)
	bRunes := []rune(b)
	distance := 0
	for i := range aRunes {
		if aRunes[i] != bRunes[i] {
			distance++
		}
	}
	return distance
}
