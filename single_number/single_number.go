package single_number

func singleNumber(nums []int) int {
	ret := 0
	for i := range nums {
		ret ^= nums[i]
	}
	return ret
}
