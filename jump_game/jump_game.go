package jump_game

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	step := nums[0]
	for i := 1; i < len(nums); i++ {
		if step > 0 {
			step--
			step = max(step, nums[i])
		} else {
			return false
		}
	}
	return true
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
