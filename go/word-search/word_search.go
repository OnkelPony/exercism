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

func findWord(word string, puzzle []string) [2][2]int{
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
		
	return [2][2]int{{-1, -1}, {-1, -1}} 
}
