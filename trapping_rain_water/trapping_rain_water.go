package trapping_rain_water

func trap(height []int) int {
	left, right, secHeight, area := 0, len(height)-1, 0, 0
	for left < right {
		if height[left] < height[right] {
			secHeight = max(secHeight, height[left])
			area += secHeight - height[left]
			left++
		} else {
			secHeight = max(secHeight, height[right])
			area += secHeight - height[right]
			right--
		}
	}
	return area
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
