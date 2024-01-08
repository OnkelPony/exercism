package wordsearch

import "errors"

type Position [2]int
var invalidPosition = Position{-1, -1}

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
	initialPositions := findInitialPositions(puzzle, word)
	for _, startPosition := range initialPositions {
		endPosition := scanDirections(word, puzzle, startPosition)
		if endPosition != invalidPosition {
			return [2][2]int{{startPosition[0], startPosition[1]}, {endPosition[0], endPosition[1]}}
		}
	}
	return [2][2]int{{-1, -1}, {-1, -1}}
}

func findInitialPositions(puzzle []string, word string) []Position {
	initialPositions := []Position{}
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[0]); j++ {
				initialPositions = append(initialPositions, Position{j, i})
		}
	}
	return initialPositions
}

func scanDirections(word string, puzzle []string, position Position) Position {
	var found bool
	var x, y int
	directions := []Position{{0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}}
	for _, direction := range directions {
		found = true
		for i, v := range word {
			x = position[0] + i*direction[0]
			y = position[1] + i*direction[1]
			if x < 0 || y < 0 || x > len(puzzle[0])-1 || y > len(puzzle)-1 || v != rune(puzzle[y][x]) {
				found = false
				break
			}
		}
		if found {
			return Position{x, y}
		}
	}
	return invalidPosition
}
