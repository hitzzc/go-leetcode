package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	left := 1
	right := 1
	for i := 0; i < len(nums); i++ {
		if i <= (len(nums)-1)/2 {
			ret[i] = 1
			ret[len(nums)-i-1] = 1
		}
		ret[len(nums)-i-1] *= right
		ret[i] *= left
		left *= nums[i]
		right *= nums[len(nums)-i-1]
	}
	return ret
}
