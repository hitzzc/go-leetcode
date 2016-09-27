package largest_rectangle_in_histogram

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	heights = append(heights, 0)
	maxArea := 0
	stack := []int{0}
	for i := 1; i < len(heights); i++ {
		if heights[i] < heights[stack[len(stack)-1]] {
			j := len(stack) - 1
			for ; j >= 0; j-- {
				if heights[stack[j]] < heights[i] {
					break
				}
				var width int
				if j == 0 {
					width = i
				} else {
					width = i - stack[j-1] - 1
				}
				area := width * heights[stack[j]]
				if area > maxArea {
					maxArea = area
				}
			}
			stack = stack[:j+1]
		}
		stack = append(stack, i)
	}
	return maxArea
}
