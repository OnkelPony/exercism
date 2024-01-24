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

func getColMins(m *Matrix) []Pair {
	result := make([]Pair, len((*m)[0]))
	cols := m.Cols()
	for i := range cols {
		minValue := cols[i][0]
		for j := range cols[i] {
			if cols[i][j] <= minValue {
				minValue = cols[i][j]
				result[i] = Pair{j+1, i+1}
			}
		}
	}
	return result
}

func getRowMaxes(m *Matrix) []Pair {
	result := make([]Pair, len((*m)))
	rows := m.Rows()
	for j := range rows {
		maxValue := rows[j][0]
		for i := range rows[j] {
			if rows[j][i] >= maxValue {
				maxValue = rows[j][i]
				result[j] = Pair{j+1, i+1}
			}
		}
	}
	return result
}
