package matrix

import "slices"

// Define the Matrix and Pair types here.
type Pair [2]int

func (m *Matrix) Saddle() []Pair {
var result []Pair
rowMaxes := getRowMaxes(m)
colMins := getColMins(m)
for _, rowMax := range rowMaxes {
	if slices.Contains(colMins, rowMax) {
		result = append(result, rowMax)
	}
}
return result
}

func getColMins(m *Matrix) []Pair{
	panic("unimplemented")
}

func getRowMaxes(m *Matrix) []Pair{
	panic("unimplemented")
}
