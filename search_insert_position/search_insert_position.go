package search_insert_position

func searchInsert(nums []int, target int) int {
	return searchInsertI(nums, 0, len(nums)-1, target)
}

func searchInsertI(nums []int, left, right, target int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if left > right {
		return -1
	}
	if nums[left] > target {
		return left
	}
	if nums[right] < target {
		return right + 1
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return searchInsertI(nums, mid+1, right, target)
	}
	return searchInsertI(nums, left, mid-1, target)
}
