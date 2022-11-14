// Package raindrops implements method, which returns string of raindrop sound, based on factor of number argument.
package raindrops

import "strconv"

const (
	divBy3 = "Pling"
	divBy5 = "Plang"
	divBy7 = "Plong"
)

// Convert returns strings based on divisibility of argument by 3, 5 and 7
func Convert(input int) string {
	result := ""
	if input%3 == 0 {
		result += divBy3
	}
	if input%5 == 0 {
		result += divBy5
	}
	if input%7 == 0 {
		result += divBy7
	}
	if result == "" {
		return strconv.Itoa(input)
	}
	return result
}
