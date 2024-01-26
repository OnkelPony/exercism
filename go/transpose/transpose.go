package transpose

func Transpose(input []string) []string {
	if len(input) == 0 {
		return input
	}
	h := 0 
	// w := len(input)
	for _, row := range input {
		if len(row) > h {
			h = len(row)
		}
	}
	result := make([]string, h)
	return result

}
