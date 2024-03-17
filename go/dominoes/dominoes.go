package dominoes

// Define the Domino type here.
type Domino [2]int

func MakeChain(input []Domino) ([]Domino, bool) {
	if len(input) == 0 || len(input) == 1 && input[0][0] == input[0][1] {
		return input, true
	}
	if len(input) == 1 {
		return nil, false
	}
	result := make([]Domino, len(input))
	result[0] = input[0]
	//Stone con't be key of the map, stones can be non-unique
	unusedStones := makeUnusedStones(input)
	for i := 1; i < len(input); i++ {
		nextStone, err := pickNext(unusedStones, result[i-1])
		if err != nil {
			return nil, false
		}
		result[i] = nextStone
	}
	return result, true
}

func pickNext(unusedStones map[int]Domino, domino Domino) (Domino, error) {
	return domino, nil
}

func makeUnusedStones(input []Domino) map[int]Domino {
	result := make(map[int]Domino, len(input)-1)
	for i := 1; i < len(input); i++ {
		result[i-1] = input[i]
	}
	return result
}
