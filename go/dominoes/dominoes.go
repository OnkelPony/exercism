package dominoes

import "errors"

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

func pickNext(unusedStones map[int]Domino, last Domino) (Domino, error) {
	//TODO: Implement "Candidates" for next stone to try more stones matching
	for k, v := range unusedStones {
		if last[1] == v[0] {
			delete(unusedStones, k)
			return v, nil
		} else if last[1] == v[1] {
			delete(unusedStones, k)
			return Domino{v[1], v[0]}, nil
		}
	}
	return Domino{}, errors.New("no matching Domino")
}

func makeUnusedStones(input []Domino) map[int]Domino {
	result := make(map[int]Domino, len(input)-1)
	for i := 1; i < len(input); i++ {
		result[i-1] = input[i]
	}
	return result
}
