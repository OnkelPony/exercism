/*
Package bob implements answers of a teenager, depending on string in parameter
*/
package bob

import "strings"

// Returns an answer depending on particular qualities of the parameter
func Hey(remark string) string {
	if isSilence(remark){
		return "Fine. Be that way!"
	}

	if isAQuestion(remark) && isAllCapitals(remark) {
		return "Calm down, I know what I'm doing!"
	}

	if isAQuestion(remark) {
		return "Sure."
	}

	if isAllCapitals(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func isSilence(remark string) bool{
	if strings.TrimSpace(remark) == "" {
		return true
	}
	return false
}

func isAllCapitals(remark string) bool {
	if strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark {
		return true
	}
	return false
}

func isAQuestion(remark string) bool {
	remark = strings.TrimSpace(remark)
	if remark[len(remark) - 1] == '?' {
		return true
	}
	return false
}
