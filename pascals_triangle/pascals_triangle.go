package pascals_triangle

func generate(numRows int) [][]int {
	rets := make([][]int, numRows)
	if numRows == 0 {
		return rets
	}
	rets[0] = []int{1}
	for i := 1; i < numRows; i++ {
		rets[i] = make([]int, i+1)
		rets[i][0] = 1
		rets[i][i] = 1
		for j := 1; j < i; j++ {
			rets[i][j] = rets[i-1][j-1] + rets[i-1][j]
		}
	}
	return rets
}
