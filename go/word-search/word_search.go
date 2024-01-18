package wordsearch

import "errors"


var invalidPosition = [2]int{-1, -1}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	var err error
	for _, word := range words {
		result[word] = findWord(word, puzzle)
		if result[word] == [2][2]int{invalidPosition, invalidPosition} {
			err = errors.New("word not found")
			return result, err
		}
	}
	return result, nil
}

func findWord(word string, puzzle []string) [2][2]int {
	for y := 0; y < len(puzzle); y++ {
		for x := 0; x < len(puzzle[0]); x++ {
			endPosition := scanDirections(word, puzzle, [2]int{x, y})
			if endPosition != invalidPosition {
				return [2][2]int{{x, y}, endPosition}
			}
		}
	}
	return [2][2]int{invalidPosition, invalidPosition}
}

func scanDirections(word string, puzzle []string, position [2]int) [2]int {
	if word[0] != puzzle[position[1]][position[0]] {
		return invalidPosition
	}
	var x, y int
	directions := [][2]int{{0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}}
	for _, direction := range directions {
		for i, _ := range word {
			x = position[0] + i*direction[0]
			y = position[1] + i*direction[1]
			if x < 0 || y < 0 || x >= len(puzzle[0]) || y >= len(puzzle) || word[i] != puzzle[y][x] {
				break
			}
			if i == len(word)-1 {
				return [2]int{x, y}
			}
		}
	}
	return invalidPosition
}
