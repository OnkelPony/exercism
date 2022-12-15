package piglatin

import "strings"

var vowels = "aeiou"
var consonants = "bcdfghjklmnpqrstvwxyz"

// Sentence returns string in parameter as Pig Latin string.
func Sentence(stc string) string {
	words := strings.Split(stc, " ")
	var result []string
	for _, input := range words {
		fh := ""
		sh := ""
		// Cases when word starts with vowels
		if strings.Contains(vowels, string(input[0])) || input[:2] == "yt" || input[:2] == "xr" {
			return input + "ay"
		}
		for i, l := range input {
			if string(input[i]) == "q" && string(input[i+1]) == "u" {
				sh += "qu"
				fh = input[i+2:]
				break
			}
			// After consonants there is vowel or 'y'
			if i > 0 && string(l) == "y" || strings.Contains(vowels, string(l)) {
				fh = input[i:]
				break
			} else {
				sh += string(l)
			}
		}
		result = append(result, fh+sh+"ay")
	}
	return strings.Join(result, " ")
}
