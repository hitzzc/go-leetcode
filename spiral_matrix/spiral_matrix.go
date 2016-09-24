package spiral_matrix

const (
	up    string = "up"
	down         = "down"
	left         = "left"
	right        = "right"
)

func spiralOrder(matrix [][]int) []int {
	ret := []int{}
	if len(matrix) == 0 {
		return ret
	}
	rows, cols := len(matrix), len(matrix[0])
	status := right
	i, j := 0, -1
	for rows != 0 && cols != 0 {
		switch status {
		case right:
			cnt := cols
			for {
				if cnt == 0 {
					break
				}
				j, cnt = j+1, cnt-1
				ret = append(ret, matrix[i][j])
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
				ret = append(ret, matrix[i][j])
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
				ret = append(ret, matrix[i][j])
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
				ret = append(ret, matrix[i][j])
			}
			cols--
			status = right
		}
	}
	return ret
}
