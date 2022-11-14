// Package pangram implements method, which decides, if entered string is pangram.
package pangram

import "unicode"

// IsPangram returns true, if string in argument contains all ascii letters.
func IsPangram(input string) interface{} {
	asciiLetters := alphabet()
	for _, letter := range input {
		letter = unicode.ToLower(letter)
		delete(asciiLetters, byte(letter))
	}
	if len(asciiLetters) == 0 {
		return true
	}
	return false
}

func alphabet() map[byte]byte {
	const letterCount = 26
	p := make(map[byte]byte, letterCount)
	for i := 0; i < letterCount; i++ {
		p['a'+byte(i)] = 0
	}
	return p
}
