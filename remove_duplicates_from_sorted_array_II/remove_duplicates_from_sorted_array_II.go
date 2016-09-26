package remove_duplicates_from_sorted_array_II

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	begin := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[begin] || begin == 0 || nums[i] != nums[begin-1] {
			begin++
			nums[begin], nums[i] = nums[i], nums[begin]
		}
	}
	return begin + 1
}
