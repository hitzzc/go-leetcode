package single_number_II

func singleNumber(nums []int) int {
	var one, two, three int
	for i := range nums {
		two |= one & nums[i]
		one ^= nums[i]
		three = one & two
		one &= ^three
		two &= ^three
	}
	return one
}
