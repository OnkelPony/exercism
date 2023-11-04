package diamond

import (
	"bytes"
	"fmt"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("invalid input character")
	}
	size := int(char-'A')*2 + 1
	var buf bytes.Buffer
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			diff := size/2 - abs(i-size/2)
			if abs(j-size/2) == diff {
				buf.WriteByte(byte('A' + diff))
			} else {
				buf.WriteByte(' ')
			}
		}
		if i < size-1 {
			buf.WriteByte('\n')
		}
	}
	return buf.String(), nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
