package house

import (
	"fmt"
	"strings"
)

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
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("This is the %s", verses[v].S))
	if v > 1 {
		builder.WriteString("\n")
	}
	for i := v; i > 1; i-- {
		builder.WriteString(fmt.Sprintf("that %s the %s", verses[i].D, verses[i-1].S))
		if i > 2 {
			builder.WriteString("\n")
		}
	}
	builder.WriteString(".")
	return builder.String()
}


func Song() string {
	var builder strings.Builder
	for i := 1; i <= len(verses); i++ {
		builder.WriteString(Verse(i))
		if i < len(verses) {
			builder.WriteString("\n\n")
		}
	}
	return builder.String()
}
