package minimum_size_subarray_sum

func minSubArrayLen(s int, nums []int) int {
	var left, right, res int
	min := len(nums) + 1
	for right < len(nums) {
		for ; res < s && right < len(nums); right++ {
			res += nums[right]
		}
		for ; res >= s; left++ {
			res -= nums[left]
			if right-left < min {
				min = right - left
			}
		}
	}
	if min == len(nums)+1 {
		return 0
	}
	return min
}
