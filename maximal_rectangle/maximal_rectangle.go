package maximal_rectangle

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	max := 0
	for i := 0; i < len(matrix); i++ {
		heights := make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				for k := i; k < len(matrix); k++ {
					if matrix[k][j] == '0' {
						break
					}
					heights[j]++
				}
			}
		}
		area := maxArea(heights)
		if area > max {
			max = area
		}
	}
	return max
}

func maxArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	max := 0
	stack := []int{0}
	heights = append(heights, 0)
	for i := 1; i < len(heights); i++ {
		if heights[i] < heights[stack[len(stack)-1]] {
			j := len(stack) - 1
			for ; j >= 0; j-- {
				if heights[stack[j]] < heights[i] {
					break
				}
				width := i
				if j != 0 {
					width = i - stack[j-1] - 1
				}
				area := width * heights[stack[j]]
				if area > max {
					max = area
				}
			}
			stack = stack[:j+1]
		}
		stack = append(stack, i)
	}
	return max
}
