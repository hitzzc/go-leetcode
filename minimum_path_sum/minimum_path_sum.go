package minimum_path_sum

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				matrix[i][j] = grid[0][0]
			} else if i == 0 {
				matrix[i][j] = grid[i][j] + matrix[i][j-1]
			} else if j == 0 {
				matrix[i][j] = grid[i][j] + matrix[i-1][j]
			} else {
				matrix[i][j] = matrix[i][j-1]
				if matrix[i][j] > matrix[i-1][j] {
					matrix[i][j] = matrix[i-1][j]
				}
				matrix[i][j] += grid[i][j]
			}
		}
	}
	return matrix[m-1][n-1]
}
