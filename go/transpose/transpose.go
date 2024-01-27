package transpose

import "strings"

func Transpose(input []string) []string {
	if len(input) == 0 {
		return input
	}
	h := 0
	w := len(input)
	for _, row := range input {
		if len(row) > h {
			h = len(row)
		}
	}
	matrix := make([][]byte, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([]byte, w)
		for j := 0; j < w; j++ {
			if i < len(input[j]) {
				matrix[i][j] = input[j][i]
			}
		}
	}
	result := make([]string, h)
	for i := 0; i < h; i++ {
		result[i] = strings.TrimRight(string(matrix[i]), "\x00")
		result[i] = strings.ReplaceAll(result[i], "\x00", " ")
	}
	return result
}
