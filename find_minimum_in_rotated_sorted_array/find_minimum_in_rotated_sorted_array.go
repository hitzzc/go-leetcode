package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, i, j int) int {
	if i == j {
		return nums[i]
	}
	if nums[i] < nums[j] {
		return nums[i]
	}
	mid := i + (j-i)/2
	if i == mid {
		return helper(nums, i+1, j)
	}
	if nums[i] < nums[mid] {
		return helper(nums, mid+1, j)
	}
	return helper(nums, i, mid)
}
