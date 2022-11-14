// Package luhn implements method, which checks for Luhn algorithm
package luhn

// Valid returns true, if argument meets Luhn algorithm
func Valid(input string) bool {
	finalIdx := len(input) - 1
	var result, index int
	for i := finalIdx; i >= 0; i-- {
		char := input[i]
		if char == ' ' {
			continue
		}
		if char < '0' || char > '9' {
			return false
		}
		if index%2 == 0 {
			result += int(char - '0')
		} else {
			result += double(int(char - '0'))
		}
		index++
	}
	return result%10 == 0 && index > 1
}

// Double returns doubled argument. If result exceeds 9, then result is lowered by 9.
func double(input int) int {
	result := 2 * input
	if result > 9 {
		result -= 9
	}
	return result
}
