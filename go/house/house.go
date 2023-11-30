package house

import "fmt"

type SubjectDoing struct {
	S string
	D string
}

var verses = map[int]SubjectDoing{
	2: {S: "malt", D: "lay"},
	3: {S: "rat", D: "ate"},
	4: {S: "cat", D: "killed"},
}

func Verse(v int) string {
	var result string
	if v == 1 {
		return "This is the house that Jack built."
	}
	result = fmt.Sprintf("This is the %s\n", verses[v].S)
	for i := v; i > 2; i-- {
		result += fmt.Sprintf("that %s the %s\n", verses[i].D, verses[i-1].S)
	}
	result += "that lay in the house that Jack built."
	return result
}

func Song() string {
	panic("Please implement the Song function")
}
