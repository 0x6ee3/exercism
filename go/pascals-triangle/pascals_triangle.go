package pascal

// Triangle ...
func Triangle(n int) [][]int {
	values := make([][]int, n)

	for i := 0; i < n; i++ {
		values[i] = make([]int, i+1)

		values[i][0] = 1
		values[i][len(values[i])-1] = 1

		for j := 1; j < i; j++ {
			values[i][j] = values[i-1][j-1] + values[i-1][j]
		}
	}

	return values
}
