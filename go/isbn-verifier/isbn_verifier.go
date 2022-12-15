// Package isbn implements method, which checks validity of ISBN code.
package isbn

import (
	"strconv"
	"strings"
)

const ten = 10

// IsValidISBN returns true, if argument is valid ISBN code.
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != ten {
		return false
	}
	sumIsbn := 0
	for i, digit := range isbn {
		intDigit, err := strconv.Atoi(string(digit))
		if err != nil {
			if digit == 'X' && i == 9 {
				intDigit = ten
			} else {
				return false
			}
		}
		sumIsbn += (ten - i) * intDigit
	}
	return sumIsbn%11 == 0
}
