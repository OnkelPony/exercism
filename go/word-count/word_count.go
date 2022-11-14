// Package wordcount implements method, which counts the occurrence of words in sentence.
package wordcount

import "strings"

// Frequency is occurrence of particular words
type Frequency map[string]int

// WordCount returns map of words and their count in sentence, entered as argument.
func WordCount(input string) Frequency {
	input = strings.ToLower(input)
	var result = Frequency{}
	words := strings.FieldsFunc(input, func(r rune) bool {
		return strings.ContainsRune("\n .,:!&@$%^", r)
	})
	for _, word := range words {
		word = strings.Trim(word, "'")
		result[word]++
	}
	return result
}
