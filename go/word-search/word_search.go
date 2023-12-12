package wordsearch

import "errors"

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
	initialPositions := [][2]int{}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if puzzle[i][j] == word[0] {
				initialPositions = append(initialPositions, [2]int{i, j})
			}
		}
	}
	for _, position := range initialPositions {
		endPosition := scanEast(word, puzzle, position)
		if endPosition != [2]int{-1, -1} {
			return [2][2]int{position, endPosition}
		}
	}
	return [2][2]int{{-1, -1}, {-1, -1}}
}

func scanEast(word string, puzzle []string, position [2]int) [2]int {
	if len(word) <= len(puzzle[0])-position[1] {
		for i, v := range word {
			if v != rune(puzzle[position[0]][position[1]+i]) {
				return [2]int{-1, -1}
			}
		}
		return [2]int{position[1] + len(word) - 1, position[0]}
	}
	return [2]int{-1, -1}
}
