package find_minimum_in_rotated_sorted_array_II

func findMin(nums []int) int {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	if nums[left] < nums[right] {
		return nums[left]
	}
	mid := left + (right-left)/2
	if mid == left {
		return helper(nums, left+1, right)
	}
	if nums[left] == nums[mid] && nums[mid] == nums[right] {
		for ; left < right && nums[left] == nums[left+1]; left++ {
		}
		for ; left < right && nums[right] == nums[right-1]; right-- {
		}
		return helper(nums, left, right)
	}
	if nums[left] <= nums[mid] && nums[mid] > nums[right] {
		return helper(nums, mid+1, right)
	}
	return helper(nums, left, mid)
}
