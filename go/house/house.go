package house

import "fmt"

type SubjectDoing struct {
	S string
	D string
}

var verses = map[int]SubjectDoing{
	1:  {S: "house that Jack built", D: ""},
	2:  {S: "malt", D: "lay in"},
	3:  {S: "rat", D: "ate"},
	4:  {S: "cat", D: "killed"},
	5:  {S: "dog", D: "worried"},
	6:  {S: "cow with the crumpled horn", D: "tossed"},
	7:  {S: "maiden all forlorn", D: "milked"},
	8:  {S: "man all tattered and torn", D: "kissed"},
	9:  {S: "priest all shaven and shorn", D: "married"},
	10: {S: "rooster that crowed in the morn", D: "woke"},
	11: {S: "farmer sowing his corn", D: "kept"},
	12: {S: "horse and the hound and the horn", D: "belonged to"},
}

func Verse(v int) string {
	var result string
	result = fmt.Sprintf("This is the %s", verses[v].S)
	if v > 1 {
		result += "\n"
	}
	for i := v; i > 1; i-- {
		result += fmt.Sprintf("that %s the %s", verses[i].D, verses[i-1].S)
		if i > 2 {
			result += "\n"
		}
	}
	result += "."
	return result
}

func Song() string {
	var result string
	for i := 1; i <= len(verses); i++ {
		result += Verse(i)
		if i < len(verses) {
			result += "\n\n"
		}
	}
	return result
}
