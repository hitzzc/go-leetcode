package maximum_product_subarray

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	ret, preMax, preMin := nums[0], nums[0], nums[0]
	var temp int
	for i := 1; i < len(nums); i++ {
		temp = preMax
		preMax = max(max(nums[i], nums[i]*preMax), nums[i]*preMin)
		preMin = min(min(nums[i], nums[i]*preMin), nums[i]*temp)
		ret = max(ret, preMax)
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
