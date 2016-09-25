package sort_colors

func sortColors(nums []int) {
	r, b := 0, len(nums)-1
	for i := 0; i <= b; i++ {
		if nums[i] == 0 {
			nums[i], nums[r] = nums[r], nums[i]
			r++
		} else if nums[i] == 2 {
			nums[i], nums[b] = nums[b], nums[i]
			b--
			i--
		}
	}
	return
}
