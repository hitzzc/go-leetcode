package search_for_a_range

func searchRange(nums []int, target int) []int {
	length := len(nums)
	if length == 0 || length == 1 && nums[0] != target {
		return []int{-1, -1}
	}
	return searchRangeI(nums, 0, length-1, target)
}

func searchRangeI(nums []int, left, right, target int) (ret []int) {
	length := len(nums)
	if length == 0 || left > right {
		return []int{-1, -1}
	}
	if left == right {
		if nums[left] == target {
			return []int{left, left}
		}
		return []int{-1, -1}
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		lower := findLower(nums, left, mid, target)
		upper := findUpper(nums, mid, right, target)
		return []int{lower, upper}
	} else if nums[mid] < target {
		return searchRangeI(nums, mid+1, right, target)
	}
	return searchRangeI(nums, left, mid-1, target)
}

func findLower(nums []int, left, right, target int) int {
	length := len(nums)
	if length == 0 || left > right {
		return -1
	}
	if left == right {
		if nums[right] == target {
			return right
		}
		return -1
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		if mid > left && nums[mid-1] == target {
			return findLower(nums, left, mid-1, target)
		}
		return mid
	}
	return findLower(nums, mid+1, right, target)
}
func findUpper(nums []int, left, right, target int) int {
	length := len(nums)
	if length == 0 || left > right {
		return -1
	}
	if left == right {
		if nums[left] == target {
			return left
		}
		return -1
	}
	mid := left + (right-left)/2
	if nums[mid] == target {
		if mid < right && nums[mid+1] == target {
			return findUpper(nums, mid+1, right, target)
		}
		return mid
	}
	return findUpper(nums, left, mid-1, target)
}
