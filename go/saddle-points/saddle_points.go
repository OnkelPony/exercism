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
	var result []Pair
	cols := m.Cols()
	for i := range cols {
		minValue := cols[i][0]
		for j := range cols[i] {
			if cols[i][j] < minValue {
				result[i] = Pair{i, j}
			}
		}
	}
	return result
}

func getRowMaxes(m *Matrix) []Pair {
	var result []Pair
	rows := m.Rows()
	for j := range rows {
		maxValue := rows[j][0]
		for i := range rows[j] {
			if rows[j][i] > maxValue {
				result[j] = Pair{i, j}
			}
		}
	}
	return result
}
