package wordsearch

import "errors"

type Position [2]int

var invalidPosition = Position{-1, -1}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	var err error
	for _, word := range words {
		result[word] = findWord(word, puzzle)
		if result[word] == [2][2]int{invalidPosition, invalidPosition} {
			err = errors.New("word not found")
		}
	}
	return result, err
}

func findWord(word string, puzzle []string) [2][2]int {
	for y := 0; y < len(puzzle); y++ {
		for x := 0; x < len(puzzle[0]); x++ {
			endPosition := scanDirections(word, puzzle, Position{x, y})
			if endPosition != invalidPosition {
				return [2][2]int{Position{x, y}, {endPosition[0], endPosition[1]}}
			}
		}
	}
	return [2][2]int{invalidPosition, invalidPosition}
}

func scanDirections(word string, puzzle []string, position Position) Position {
	var x, y int
	directions := []Position{{0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}}
	for _, direction := range directions {
		for i, v := range word {
			x = position[0] + i*direction[0]
			y = position[1] + i*direction[1]
			if x < 0 || y < 0 || x > len(puzzle[0])-1 || y > len(puzzle)-1 || v != rune(puzzle[y][x]) {
				break
			}
			if i == len(word)-1 {
				return Position{x, y}
			}
		}
	}
	return invalidPosition
}
