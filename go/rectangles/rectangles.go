package rectangles

func Count(diagram []string) int {
	var result int
	for h := 0; h < len(diagram); h++ {
		for w := 0; w < len(diagram[h]); w++ {
			if diagram[h][w] == '+' {
				result += 1
			}
		}
	}
	return result
}
