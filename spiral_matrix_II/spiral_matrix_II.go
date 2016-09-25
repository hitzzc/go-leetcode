package spiral_matrix_II

const (
	up    string = "up"
	down         = "down"
	left         = "left"
	right        = "right"
)

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	if n == 0 {
		return matrix
	}
	rows, cols := n, n
	status := right
	i, j := 0, -1
	current := 1
	for rows != 0 && cols != 0 {
		switch status {
		case right:
			cnt := cols
			for {
				if cnt == 0 {
					break
				}
				j, cnt = j+1, cnt-1
				matrix[i][j] = current
				current++
			}
			rows--
			status = down
		case down:
			cnt := rows
			for {
				if cnt == 0 {
					break
				}
				i, cnt = i+1, cnt-1
				matrix[i][j] = current
				current++
			}
			cols--
			status = left
		case left:
			cnt := cols
			for {
				if cnt == 0 {
					break
				}
				j, cnt = j-1, cnt-1
				matrix[i][j] = current
				current++
			}
			rows--
			status = up
		case up:
			cnt := rows
			for {
				if cnt == 0 {
					break
				}
				i, cnt = i-1, cnt-1
				matrix[i][j] = current
				current++
			}
			cols--
			status = right
		}
	}
	return matrix
}
