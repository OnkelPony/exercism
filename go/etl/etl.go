// Package etl implements method, which transforms old style scrabble values to new ones.
package etl

import "strings"

// Transform takes scrabble values in map[int][]string and returns these values as map[string]int with string in lowercase.
func Transform(oldScrab map[int][]string) map[string]int {
	result := map[string]int{}
	for value, letterSlice := range oldScrab {
		for _, letter := range letterSlice {
			result[strings.ToLower(letter)] = value
		}
	}
	return result
}

