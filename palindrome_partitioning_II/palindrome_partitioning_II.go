package palindrome_partitioning_II

func minCut(s string) int {
	if len(s) < 2 {
		return 0
	}
	matrix := make([][]bool, len(s))
	for i := range matrix {
		matrix[i] = make([]bool, len(s))
	}
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		for j := 0; j <= i; j++ {
			if i == j || bytes[j] == bytes[i] && (i == j+1 || matrix[j+1][i-1]) {
				matrix[j][i] = true
			}
		}
	}

	records := make([]int, len(bytes))
	for i := 1; i < len(bytes); i++ {
		min := records[i-1] + 1
		for j := i - 1; j >= 0; j-- {
			if matrix[j][i] {
				if j == 0 {
					min = 0
				} else {
					record := records[j-1] + 1
					if record < min {
						min = record
					}
				}
			}
		}
		records[i] = min
	}
	return records[len(bytes)-1]
}

func isPartition(bytes []byte, i, j int) bool {
	for ; i < j; i, j = i+1, j-1 {
		if bytes[i] != bytes[j] {
			return false
		}
	}
	return true
}
