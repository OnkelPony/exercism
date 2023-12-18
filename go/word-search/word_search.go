package wordsearch

import "errors"

type Position struct {
	x int
	y int
}

var invalidPosition = Position{x: -1, y: -1}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	var err error
	for _, word := range words {
		result[word] = findWord(word, puzzle)
		if result[word] == [2][2]int{{-1, -1}, {-1, -1}} {
			err = errors.New("word not found")
		}
	}
	return result, err
}

func findWord(word string, puzzle []string) [2][2]int {
	width := len(puzzle[0])
	height := len(puzzle)
	initialPositions := []Position{}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if puzzle[i][j] == word[0] {
				initialPositions = append(initialPositions, Position{x: j, y: i})
			}
		}
	}
	for _, startPosition := range initialPositions {
		endPosition := scanEast(word, puzzle, startPosition)
		if endPosition != invalidPosition {
			return [2][2]int{{startPosition.x, startPosition.y}, {endPosition.x, endPosition.y}}
		}
	}
	return [2][2]int{{-1, -1}, {-1, -1}}
}

func scanEast(word string, puzzle []string, position Position) Position {
	var result Position
	var x, y int
	directions := []Position{{0, 1},{1, 0}}
	for _, direction := range directions {
		for i, v := range word {
			x = position.x + i*direction.x
			y = position.y + i*direction.y
			if x < 0 || y < 0 || x > len(puzzle[0])-1 || y > len(puzzle)-1 || v != rune(puzzle[y][x]) {
				result = invalidPosition
				break
			}
		}
		if result != invalidPosition {
			return Position{x, y} // adjust this to work with all directions
		}
	}
	return result
}
