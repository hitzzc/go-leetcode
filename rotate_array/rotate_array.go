package rotate_array

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
	return
}

func reverse(nums []int, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return
}
