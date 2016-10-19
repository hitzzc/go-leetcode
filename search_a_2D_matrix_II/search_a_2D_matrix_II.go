package search_a_2D_matrix_II

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for row <= len(matrix)-1 && col >= 0 {
		if matrix[row][col] == target {
			return true
		}
		for row <= len(matrix)-1 && col >= 0 && matrix[row][col] < target {
			row++
		}
		for row <= len(matrix)-1 && col >= 0 && matrix[row][col] > target {
			col--
		}
	}
	return false
}
