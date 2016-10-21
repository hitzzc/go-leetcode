package missing_number

func missingNumber(nums []int) int {
	ret := 0
	for i := range nums {
		ret ^= nums[i] ^ (i + 1)
	}
	return ret
}

func missingNumberI(nums []int) int {
	ret := 0
	for i := range nums {
		ret += i + 1 - nums[i]
	}
	return ret
}
