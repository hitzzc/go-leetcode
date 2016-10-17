package maximal_square

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m := map[byte]int{
		'0': 0,
		'1': 1,
	}
	p := make([][]int, len(matrix))
	for i := range p {
		p[i] = make([]int, len(matrix[i]))
	}
	max := 0
	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 || j == 0 {
				p[i][j] = m[matrix[i][j]]
			} else if matrix[i][j] == '1' {
				p[i][j] = min(min(p[i-1][j], p[i][j-1]), p[i-1][j-1]) + 1
			}
			if p[i][j] > max {
				max = p[i][j]
			}
		}
	}
	return max * max
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
