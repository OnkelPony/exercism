// Package encode implement methods, which use RLE to encode and decode text.
package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode returns text in argument, encoded using RLE
func RunLengthEncode(input string) string {
	result := ""
	if input == "" {
		return result
	}
	count := 1
	currentChar := input[0]
	for i, char := range input {
		if i == 0 {
			continue
		}
		if uint8(char) == input[i-1] {
			count++
		} else {
			result += numChars(count, currentChar)
			count = 1
		}
		currentChar = uint8(char)
	}
	result += numChars(count, currentChar)
	return result
}

// numChars returns number, followed by character as string. In case number is 1, it is omitted.
func numChars(count int, currentChar uint8) string {
	if count == 1 {
		return string(currentChar)
	}
	return strconv.Itoa(count) + string(currentChar)
}

// RunLengthDecode returns text in argument, decoded using RLE
func RunLengthDecode(input string) string {
	result := ""
	countString := ""
	if input == "" {
		return result
	}
	for _, char := range input {
		if unicode.IsDigit(char) {
			countString += string(char)
		} else {
			result += multiplyChar(countString, char)
			countString = ""
		}
	}
	return result
}

// multiplyChar returns char multiplied by number in countString
func multiplyChar(countString string, char int32) string {
	if countString == "" {
		return string(char)
	}
	count, _ := strconv.Atoi(countString)
	return strings.Repeat(string(char), count)
}
