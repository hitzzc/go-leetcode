package container_with_most_water

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max, left, right := 0, 0, len(height)-1
	for left < right {
		area := minHeight(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			for left < len(height)-1 {
				left++
				if height[left] > height[left-1] {
					break
				}
			}
		} else {
			for right > 0 {
				right--
				if height[right] > height[right+1] {
					break
				}
			}
		}
	}
	return max
}

func minHeight(x, y int) int {
	if x < y {
		return x
	}
	return y
}
