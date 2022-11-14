package wordy

import (
	"strconv"
	"strings"
)

// Answer returns result of the mathematical operation entered as a string.
func Answer(question string) (int, bool) {
	var result int
	var operator string
	var wasNumeric bool

	question = strings.TrimSuffix(question, "?")
	words := strings.Split(question, " ")
	words = words[2:]
	for idx := 0; idx < len(words); idx++ {
		answer, err := strconv.Atoi(words[idx])

		// answer is not a number
		if err != nil {
			if !wasNumeric {
				return 0, false
			}
			operator = words[idx]
			if operator == "multiplied" || operator == "divided" {
				idx++
			} else if operator != "plus" && operator != "minus" {
				return 0, false
			}
			wasNumeric = false

			// answer is a number
		} else {
			if wasNumeric {
				return 0, false
			}
			switch operator {
			case "":
				result = answer
			case "plus":
				result += answer
			case "minus":
				result -= answer
			case "multiplied":
				result *= answer
			case "divided":
				result /= answer
			}
			wasNumeric = true
		}
	}
	return result, wasNumeric
}
