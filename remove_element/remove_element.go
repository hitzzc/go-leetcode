package remove_element

func removeElement(nums []int, val int) int {
	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[start] = nums[i]
		start++
	}
	return start
}
