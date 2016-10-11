package single_number_II

func singleNumber(nums []int) int {
	var one, two, three int
	for i := range nums {
		three = two & nums[i]
		two = ^three & (two | one&nums[i])
		one = ^three & (one ^ nums[i])
	}
	return one
}
