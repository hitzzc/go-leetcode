package search_in_rotated_sorted_array_II

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return true
		}
		return false
	}
	return search1(nums, target, 0, len(nums)-1) != -1
}

func search1(nums []int, target int, left, right int) int {
	if left > right {
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
		return mid
	}
	if nums[left] < nums[mid] && nums[left] <= target && target < nums[mid] {
		return binarySearch(nums, target, left, mid-1)
	}
	if nums[mid] < nums[right] && nums[mid] < target && target <= nums[right] {
		return binarySearch(nums, target, mid+1, right)
	}
	if nums[left] > nums[mid] {
		return search1(nums, target, left, mid-1)
	}
	if nums[mid] > nums[right] {
		return search1(nums, target, mid+1, right)
	}
	if nums[left] == nums[mid] && nums[mid] == nums[right] {
		for ; left < right && nums[left] == nums[left+1]; left++ {
		}
		for ; left < right && nums[right] == nums[right-1]; right-- {
		}
		return search1(nums, target, left, right)
	}
	return -1
}

func binarySearch(nums []int, target, left, right int) int {
	if left > right {
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
		return mid
	}
	if nums[left] <= target && target < nums[mid] {
		return binarySearch(nums, target, left, mid-1)
	}
	if nums[mid] < target && target <= nums[right] {
		return binarySearch(nums, target, mid+1, right)
	}
	return -1
}
