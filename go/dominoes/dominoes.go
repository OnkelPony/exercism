package dominoes


// Define the Domino type here.
type Domino [2]int

func MakeChain(input []Domino) ([]Domino, bool) {
	if len(input) == 0 || len(input) == 1 && input[0][0] == input[0][1]{
		return input, true
	}
	result := make([]Domino, 0, len(input))
	result[0] = input[0]
	//Stone con't be key of the map, stones can be non-unique
	unusedStones := makeUnusedStones(input)
	for i := 1; i < len(input); i++ {
		nextStone, err := pickNext(unusedStones, result[i - 1])
		if err != nil {
			return nil, false
		}
		result[i] = nextStone
	}
	return result, true
}
