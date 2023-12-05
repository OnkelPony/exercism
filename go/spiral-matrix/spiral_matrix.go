package spiralmatrix

func SpiralMatrix(size int) [][]int {

	result := make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
	}
	
	return result
}
