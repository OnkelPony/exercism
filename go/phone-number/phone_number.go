// Package phonenumber implements methods, which parse user input and produce phone numbers in various formats.
package phonenumber

import (
	"errors"
	"strconv"
	"unicode"
)

// Number returns phone number from argument as 10 digit string.
func Number(input string) (string, error) {
	var result string
	for _, char := range input {
		if unicode.IsDigit(char) {
			result += string(char)
		}
	}
	if len(result) == 11 && result[0] == '1' {
		result = result[1:]
	}
	if len(result) != 10 {
		return "", errors.New("phone number must be 10 digits long excluding initial 1")
	}
	firsAreaDigit, _ := strconv.Atoi(string(result[0]))
	firstExchDigit, _ := strconv.Atoi(string(result[3]))
	if firsAreaDigit < 2 || firstExchDigit < 2 {
		return "", errors.New("area and exchange codes can't start with 1 or 0")
	}
	return result, nil
}

// AreaCode takes string from argument and returns 3 digit area code.
func AreaCode(input string) (string, error) {
	result, err := Number(input)
	if err == nil {
		return result[:3], err
	}
	return "", err
}

// Format takes string from argument and returns phone number in format (NXX) NXX-XXXX
func Format(input string) (string, error) {
	result, err := Number(input)
	if err == nil {
		result = "(" + result[:3] + ") " + result[3:6] + "-" + result[6:]
		return result, err
	}
	return "", err
}
