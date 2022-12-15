/*
Package acronyms implements creation of acronyms from expressions passed as arguments.
*/
package acronym

import (
	"log"
	"regexp"
	"strings"
)

// Returns string containing capitals, which are first letters of the words of given expression
func Abbreviate(s string) string {
	// Make a Regex to say we want anything except apostrophe, letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9']+")
	if err != nil {
		log.Fatal(err)
	}
	words := reg.Split(s, -1)
	acro := ""
	for _, word := range words {
		acro += string(word[0])
	}
	acro = strings.ToUpper(acro)
	return acro
}
