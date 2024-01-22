package wordsearch

import "errors"

var directions = [][2]int{{0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}}

var ErrWordNotFound = errors.New("word not found")

// Solve should have a comment
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int, len(words))
	var err error
	for _, word := range words {
		if result[word], err = findWord(word, puzzle); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func findWord(word string, puzzle []string) ([2][2]int, error) {
	for row := 0; row < len(puzzle); row++ {
		for col := 0; col < len(puzzle[0]); col++ {
			if endPosition, err := scanDirections(word, puzzle, col, row); err == nil {
				return [2][2]int{{col, row}, endPosition}, nil
			}
		}
	}
	return [2][2]int{}, ErrWordNotFound
}

func scanDirections(word string, puzzle []string, startCol, startRow int) ([2]int, error) {
	if word[0] != puzzle[startRow][startCol] {
		return [2]int{}, ErrWordNotFound
	}

	var col, row int
direction:
	for _, direction := range directions {
		distanceToEndOfWord := len(word) - 1
		endCol := startCol + direction[0]*distanceToEndOfWord
		endRow := startRow + direction[1]*distanceToEndOfWord
		if endCol < 0 || endCol >= len(puzzle[0]) || endRow < 0 || endRow >= len(puzzle) {
			continue
		}
		for i := range word {
			col = startCol + i*direction[0]
			row = startRow + i*direction[1]
			if word[i] != puzzle[row][col] {
				continue direction
			}
		}
		return [2]int{col, row}, nil
	}
	return [2]int{}, ErrWordNotFound
}
