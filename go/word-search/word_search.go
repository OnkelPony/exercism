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
	for _, position := range initialPositions {
		endPosition := scanEast(word, puzzle, position)
		if endPosition != invalidPosition {
			return [2][2]int{{position.x, position.y}, {endPosition.x, endPosition.y}}
		}
	}
	return [2][2]int{{-1, -1}, {-1, -1}}
}

func scanEast(word string, puzzle []string, position Position) Position {
	if len(word) <= len(puzzle[0])-position.x {
		for i, v := range word {
			if v != rune(puzzle[position.y][position.x+i]) {
				return invalidPosition
			}
		}
		return Position{position.x + len(word) - 1, position.y}
	}
	return invalidPosition
}
