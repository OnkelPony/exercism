// Package rotationalcipher implements Caesar cipher, based on rotating alphabet.
package rotationalcipher

import "unicode"

// RotationalCipher returns plain text with shifted position of all letters, other characters unchanged.
func RotationalCipher(plain string, shiftKey int) string {
	var result string
	for _, char := range plain {
		if unicode.IsLetter(char) {
			result += shift(char, shiftKey)
		} else {
			result += string(char)
		}
	}
	return result
}

// shift returns character, shifted by key amount.
func shift(char int32, key int) string {
	key = key % 26
	if unicode.ToLower(char)+int32(key) > 'z' {
		char += int32(key - 26)
	} else {
		char += int32(key)
	}
	return string(char)
}
