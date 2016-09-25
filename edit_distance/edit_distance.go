package edit_distance

func minDistance(word1 string, word2 string) int {
	bytes1, bytes2 := []byte(word1), []byte(word2)
	if len(bytes1) == 0 {
		return len(bytes2)
	}
	if len(bytes2) == 0 {
		return len(bytes1)
	}
	matrix := make([][]int, len(bytes1)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(bytes2)+1)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}
	matrix[0][0] = 0

	for i := 0; i <= len(bytes1); i++ {
		for j := 0; j <= len(bytes2); j++ {
			if i > 0 {
				matrix[i][j] = min(matrix[i][j], matrix[i-1][j]+1)
			}
			if j > 0 {
				matrix[i][j] = min(matrix[i][j], matrix[i][j-1]+1)
			}
			if i > 0 && j > 0 {
				if bytes1[i-1] == bytes2[j-1] {
					matrix[i][j] = min(matrix[i][j], matrix[i-1][j-1])
				} else {
					matrix[i][j] = min(matrix[i][j], matrix[i-1][j-1]+1)
				}
			}
		}
	}
	return matrix[len(bytes1)][len(bytes2)]
}

func min(i, j int) int {
	if i == -1 {
		return j
	}
	if j == -1 {
		return i
	}
	if i < j {
		return i
	}
	return j
}
