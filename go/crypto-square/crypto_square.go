// Package cryptosquare implements method, which encodes text using square cypher
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode returns argument, encoded using crypto square coding.
func Encode(input string) string {
	result := ""
	plainText := ""
	input = strings.ToLower(input)
	for _, char := range input {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			plainText += string(char)
		}
	}
	plainTextLength := len(plainText)
	cols := int(math.Ceil(math.Sqrt(float64(plainTextLength))))
	rows := int(math.Ceil(float64(plainTextLength) / float64(cols)))
	for c := 0; c < cols; c++ {
		if c > 0 {
			result += " "
		}
		for r := 0; r < rows; r++ {
			index := c + cols*r
			if index >= plainTextLength {
				result += " "
			} else {
				result += string(plainText[index])
			}
		}
	}

	return result
}
