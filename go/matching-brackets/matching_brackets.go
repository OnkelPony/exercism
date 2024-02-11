package brackets

import "slices"

var openBrackets = []byte{'(', '[', '{'}
var closeBrackets = []byte{')', ']', '}'}
var bracketPairs = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

func Bracket(input string) bool {
	brackets := []byte{}
	for i := 0; i < len(input); i++ {
		if slices.Contains(openBrackets, input[i]) {
			brackets = append(brackets, input[i])
		} else if slices.Contains(closeBrackets, input[i]) {
			if bracketPairs[brackets[len(brackets)-1]] == input[i] {
				brackets = brackets[:len(brackets)-1]
			} else {
				return false
			}
		}
	}
	return len(brackets) == 0
}
