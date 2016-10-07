package pascals_triangle_II

func getRow(rowIndex int) []int {
	rets := make([]int, rowIndex+1)
	rets[0] = 1
	for i := 1; i <= rowIndex; i++ {
		rets[i] = 1
		for j := i - 1; j > 0; j-- {
			rets[j] += rets[j-1]
		}
	}
	return rets
}
