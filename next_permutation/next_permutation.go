package next_permutation

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	swapIndex := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			swapIndex = i - 1
			break
		}
	}
	if swapIndex == -1 {
		reverseNums(nums)
		return
	}
	for i := len(nums) - 1; i > swapIndex; i-- {
		if nums[i] > nums[swapIndex] {
			nums[i], nums[swapIndex] = nums[swapIndex], nums[i]
			reverseNums(nums[swapIndex+1:])
			break
		}
	}
	return
}

func reverseNums(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return
}
