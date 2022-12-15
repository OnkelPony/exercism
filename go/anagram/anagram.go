package anagram

import (
	"sort"
	"strings"
)

func Anagram(subject, candidate string) bool {
	if(strings.ToLower(subject) == strings.ToLower(candidate)) {
		return false
	}
	s := []rune(strings.ToLower(subject))
	sort.Slice(s, func(a, b int) bool {return s[a] < s[b]})
	subject = string(s)
	s = []rune(strings.ToLower(candidate))
	sort.Slice(s, func(a, b int) bool {return s[a] < s[b]})
	candidate = string(s)
	return subject == candidate
}

func Detect(subject string, candidates []string) []string {
	result := []string{}
	for _,candidate := range candidates {
		if Anagram(subject, candidate) {
			result = append(result, candidate)
		}
	}
	return result
}