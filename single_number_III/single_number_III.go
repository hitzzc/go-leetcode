package single_number_III

func singleNumber(nums []int) []int {
	var allox int
	for i := range nums {
		allox ^= nums[i]
	}
	mask := 1
	for ; mask&allox == 0; mask <<= 1 {
	}
	var ret1, ret2 int
	for i := range nums {
		if nums[i]&mask != 0 {
			ret1 ^= nums[i]
		} else {
			ret2 ^= nums[i]
		}
	}
	return []int{ret1, ret2}
}
