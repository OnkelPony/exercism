package rectangles

func Count(diagram []string) int {
	var result int
	for h := 0; h < len(diagram) - 1; h++ {
		for w := 0; w < len(diagram[h]) - 1; w++ {
			if diagram[h][w] == '+' {
				result += countRectangles(diagram, h, w)
			}
		}
	}
	return result
}

func countRectangles(diagram []string, h, w int) int {
	var result int
	for y := h + 1; y < len(diagram); y++ {
		for x := w + 1; x < len(diagram[h]); x++ {
			if diagram[y][x] == '+' && diagram[y][w] == '+' && diagram[h][x] == '+' {
				result++
			}
		}
	}
	return result
}
