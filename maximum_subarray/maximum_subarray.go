package maximum_subarray

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	preSum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum := 0
		if preSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum = nums[i] + preSum
		}
		if currentSum > max {
			max = currentSum
		}
		preSum = currentSum
	}
	return max
}
