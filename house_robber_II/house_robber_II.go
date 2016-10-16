package house_robber_II

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	x := helper(nums, 0, len(nums)-2)
	y := helper(nums, 1, len(nums)-1)
	return max(x, y)
}

func helper(nums []int, i, j int) int {
	var n1, n2, current int
	for ; i <= j; i++ {
		current = max(n1, n2+nums[i])
		n2 = n1
		n1 = current
	}
	return n1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
