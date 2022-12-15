// Package cipher implement several methods to handle en- and decrypting using various ciphers.
package cipher

import (
	"math"
	"unicode"
)

// NewCaesar returns Caesar cipher.
func NewCaesar() Cipher {
	return Caesar{}
}

// Caesar implements Cipher interface
type Caesar struct {
}

// Encode returns argument encoded using Casesar cipher
func (Caesar) Encode(input string) string {
	c := NewShift(3)
	return c.Encode(input)
}

// Decode returns argument decoded using Casesar cipher
func (Caesar) Decode(input string) string {
	c := NewShift(3)
	return c.Decode(input)
}

// NewShift returns Shift cipher.
func NewShift(howMany int) Cipher {
	if math.Abs(float64(howMany)) > 25 || math.Abs(float64(howMany)) < 1 {
		return nil
	}
	return Shift{howMany: howMany}
}

// Shift implements Cipher interface
type Shift struct {
	howMany int
}

// Encode returns argument encoded using Casesar cipher
func (s Shift) Encode(input string) string {
	var result string
	for _, char := range input {
		char = unicode.ToLower(char)
		if unicode.IsLetter(char) {
			result += shift(char, s.howMany)
		}
	}
	return result
}

// Decode returns argument decoded using Casesar cipher
func (s Shift) Decode(input string) string {
	var result string
	for _, char := range input {
		result += shift(char, -s.howMany)
	}
	return result
}

// NewVigenere returns Vigenere cipher
func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}
	var numOfAs int
	for _, char := range key {
		if !unicode.IsLower(char) {
			return nil
		}
		if char == 'a' {
			numOfAs++
		}
	}
	if numOfAs == len(key) {
		return nil
	}
	return Vigenere{key: key}
}

// Vigenere implements Cipher interface
type Vigenere struct {
	key string
}

// Encode returns argument encoded using Vigenere cipher
func (v Vigenere) Encode(input string) string {
	var result string
	var keyIdx int
	key := v.key
	for _, char := range input {
		char = unicode.ToLower(char)
		if unicode.IsLetter(char) {
			shiftValue := key[keyIdx%len(key)] - 'a'
			result += shift(char, int(shiftValue))
			keyIdx++
		}
	}
	return result
}

// Decode returns argument decoded using Vigenere cipher.
func (v Vigenere) Decode(input string) string {
	var result string
	var keyIdx int
	key := v.key
	for _, char := range input {
		shiftValue := key[keyIdx%len(key)] - 'a'
		result += shift(char, -int(shiftValue))
		keyIdx++
	}
	return result
}

//shift returns character from argument as string, moved by shiftValue in ASCII table.
func shift(char int32, shiftValue int) string {
	if char+int32(shiftValue) > 'z' {
		char += int32(shiftValue) - 26
	} else if char+int32(shiftValue) < 'a' {
		char += int32(shiftValue) + 26
	} else {
		char += int32(shiftValue)
	}
	return string(char)
}
