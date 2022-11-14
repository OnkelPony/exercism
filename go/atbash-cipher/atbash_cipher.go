// Package atbash implements Atbash cipher. It encodes text as a mirros, "a" becomes "z", "b" becomes "y" etc.
package atbash

import (
	"strings"
	"unicode"
)

const divideBy int = 5

// Atbash returns string from argument, encoded by Atbash cipher.
func Atbash(clearText string) string {
	var result string
	var groupCount int
	for _, char := range clearText {
		char = unicode.ToLower(char)
		if char >= 'a' && char <= 'z' {
			result += string('z' - char + 'a')
		} else if char >= '0' && char <= '9' {
			result += string(char)
		} else {
			continue
		}
		groupCount++
		if groupCount == divideBy {
			groupCount = 0
			result += " "
		}
	}
	return strings.TrimSpace(result)
}
