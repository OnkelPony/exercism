package spiralmatrix

func SpiralMatrix(size int) [][]int {

	result := make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
	}
	col, row := 0, 0
	dCol, dRow := 1, 0
	for i := 1; i <= size*size; i++ {
		result[row][col] = i
		if col+dCol == size || row+dRow == size || col+dCol < 0 || result[row+dRow][col+dCol] > 0 {
			dCol, dRow = -dRow, dCol
		}
		row += dRow
		col += dCol
	}
	return result
}
