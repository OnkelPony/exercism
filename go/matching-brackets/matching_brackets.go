package brackets

func Bracket(input string) bool {
	var stack []byte
	for i := 0; i < len(input); i++ {
		char := input[i] 
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != matchingBracket(int(char)) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func matchingBracket(char int) byte {
	switch char {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	default:
		return 0
	}
}
