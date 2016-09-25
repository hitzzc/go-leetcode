package set_matrix_zeroes

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	var clearRow, clearCol bool
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					clearRow = true
				}
				if j == 0 {
					clearCol = true
				}
				if i != 0 && j != 0 {
					matrix[0][j] = 0
					matrix[i][0] = 0
				}
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
	if clearRow {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if clearCol {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
	return
}
