package dominoes

import "sort"

// Define the Domino type here.
type Domino [2]int

func MakeChain(input []Domino) ([]Domino, bool) {
	sort.Slice(input, func(i, j int) bool {
		return input[i][1] == input[j][1] || input[i][1] == input[j][0]
	})
	return input, true
}
