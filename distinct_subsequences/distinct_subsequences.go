package distinct_subsequences

func numDistinct(s string, t string) int {
	return numDistinctBytes([]byte(s), []byte(t))
}

func numDistinctBytes(s []byte, t []byte) int {
	if len(s) == 0 {
		return 0
	}
	if len(t) == 0 {
		return 1
	}
	matrix := make([][]int, len(s)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(t)+1)
		matrix[i][0] = 1
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if s[i-1] != t[j-1] {
				matrix[i][j] = matrix[i-1][j]
			} else {
				matrix[i][j] = matrix[i-1][j-1] + matrix[i-1][j]
			}
		}
	}
	return matrix[len(s)][len(t)]
}
