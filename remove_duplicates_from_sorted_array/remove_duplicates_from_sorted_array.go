package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	start, preVal := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != preVal {
			start, preVal, nums[start] = start+1, nums[i], nums[i]
		}
	}
	return start
}
