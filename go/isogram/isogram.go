// Package isogram implements method, which finds isograms.
package isogram

import "unicode"

// IsIsogram returns true, if string argument is isogram.
func IsIsogram(input string) bool {
	lettersCount := make(map[rune]int)
	for _, letter := range input {
		letter = unicode.ToUpper(letter)
		if letter != ' ' && letter != '-' {
			lettersCount[letter]++
		}
		if lettersCount[letter] > 1 {
			return false
		}
	}
	return true
}
