package pascal

func Triangle(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, i+1)
	}
	for i := range result {
		for j := range result[i] {
			if j == 0 || j == i {
				result[i][j] = 1
			} else {
				result[i][j] = result[i - 1][j-1] + result[i -1][j]
			}
		}
	}
	return result
}
