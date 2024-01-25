package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix represents two-dimensional matrix of integers with non-zero rows and columns.
type Matrix [][]int

// New creates new matrix from string, where columns are separated by spaces and rows by newlines.
func New(s string) (*Matrix, error) {
	if s == "" {
		result := make(Matrix, 0)
		return &result, nil
	}
	lines := strings.Split(s, "\n")
	result := make(Matrix, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		elements := strings.Split(line, " ")
		result[i] = make([]int, len(elements))
		for j, element := range elements {
			intElement, err := strconv.Atoi(element)
			if err != nil {
				return nil, err
			}
			result[i][j] = intElement
		}
		if i > 0 && len(result[i]) != len(result[i-1]) {
			return nil, errors.New("uneven rows")
		}
	}
	return &result, nil
}

// Cols must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	if len(*m) == 0 {
		return [][]int{}
	}
	result := make(Matrix, len((*m)[0]))
	for i := range (*m)[0] {
		result[i] = make([]int, len(*m))
		for j := range *m {
			result[i][j] = (*m)[j][i]
		}
	}
	return result
}

// Rows must return the results without affecting the matrix.
func (m *Matrix) Rows() [][]int {
	if len(*m) == 0 {
		return [][]int{}
	}
	result := make(Matrix, len(*m))
	for i := range *m {
		result[i] = make([]int, len((*m)[i]))
		copy(result[i], (*m)[i])
	}
	return result
}

// Set places values into the matrix position col and row, returns true if successful.
func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(*m) {
		return false
	}
	if col < 0 || col >= len((*m)[0]) {
		return false
	}
	(*m)[row][col] = val
	return true
}
