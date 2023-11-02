package diamond

import (
	"bytes"
	"fmt"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("Invalid input character")
	}
	size := int(char-'A')*2 + 1
	result := make([][]byte, size)
	for i := range result {
		result[i] = make([]byte, size)
	}
	for i := range result {
		for j := range result[i] {
			if i+j == int(char-'A') {
				result[i][j] = '#'
			} else {
				result[i][j] = ' '
			}
		}
	}
	return string(bytes.Join(result, []byte("\n"))), nil
}
