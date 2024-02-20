package rectangles

type Point struct {
	x int
	y int
}

// Count counts the number of rectangles in the diagram.
// The diagram is represented as a slice of strings, where each string represents a row of the diagram.
// A '+' character represents corner, '-' and '|' characters represent sides of the rectangle.
// The function returns the total number of rectangles in the diagram.
func Count(diagram []string) int {
	var result int
	for h := 0; h < len(diagram)-1; h++ {
		for w := 0; w < len(diagram[h])-1; w++ {
			if diagram[h][w] == '+' {
				result += countRectangles(diagram, Point{w, h})
			}
		}
	}
	return result
}

func countRectangles(diagram []string, topLeft Point) int {
	var result int
	var botRight Point
	for botRight.y = topLeft.y + 1; botRight.y < len(diagram); botRight.y++ {
		for botRight.x = topLeft.x + 1; botRight.x < len(diagram[topLeft.y]); botRight.x++ {
			if diagram[botRight.y][botRight.x] == '+' &&
				diagram[botRight.y][topLeft.x] == '+' &&
				diagram[topLeft.y][botRight.x] == '+' &&
				isRectangle(diagram, topLeft, botRight) {
				result++
			}
		}
	}
	return result
}

func isRectangle(diagram []string, topLeft, botRight Point) bool {
	for x := topLeft.x + 1; x < botRight.x; x++ {
		if diagram[topLeft.y][x] != '+' &&
			diagram[topLeft.y][x] != '-' ||
			diagram[botRight.y][x] != '+' &&
				diagram[botRight.y][x] != '-' {
			return false
		}
	}
	for y := topLeft.y + 1; y < botRight.y; y++ {
		if diagram[y][topLeft.x] != '+' &&
			diagram[y][topLeft.x] != '|' ||
			diagram[y][botRight.x] != '+' &&
				diagram[y][botRight.x] != '|' {
			return false
		}
	}
	return true
}
