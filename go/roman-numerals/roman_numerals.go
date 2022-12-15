// Package romannumerals implements method, which converts indian numbers to roman ones.
package romannumerals

import "errors"

var divisors = [...]int{1000, 500, 100, 50, 10, 5, 1}
var romanSymbols = [...]string{"M", "D", "C", "L", "X", "V", "I"}

// ToRomanNumeral takes int number and returns its equivalent in roman numerals
func ToRomanNumeral(indo int) (string, error) {
	result := ""
	var err error
	if indo < 1 || indo > 3000 {
		err = errors.New("Tohle nepude!")
	}
	result += compose(indo/divisors[0], romanSymbols[0])
	indo %= divisors[0]
	for i := 1; i < len(divisors); i += 2 {
		quotient := indo / divisors[i]
		remainderBy5 := indo % divisors[i] / divisors[i+1]
		indo %= divisors[i+1]
		if quotient == 1 {
			if remainderBy5 == 4 {
				result += compose(1, romanSymbols[i+1]) + compose(1, romanSymbols[i-1])
			} else {
				result += compose(1, romanSymbols[i]) + compose(remainderBy5, romanSymbols[i+1])
			}
		} else {
			if remainderBy5 == 4 {
				result += compose(1, romanSymbols[i+1]) + compose(1, romanSymbols[i])
			} else {
				result += compose(remainderBy5, romanSymbols[i+1])
			}
		}
	}
	return result, err
}

// Compose takes count and string as arguments and returns that string repeated count times
func compose(count int, romanSymbol string) string {
	component := ""
	for i := 0; i < count; i++ {
		component += romanSymbol
	}
	return component
}
