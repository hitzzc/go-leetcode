package kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	return partition(nums, 0, len(nums)-1, k)
}

func partition(nums []int, i, j, k int) (ret int) {
	p := i
	for m := i; m < j; m++ {
		if nums[m] > nums[j] {
			nums[p], nums[m] = nums[m], nums[p]
			p++
		}
	}
	nums[p], nums[j] = nums[j], nums[p]
	if p+1 == k {
		return nums[p]
	}
	if k < p+1 {
		return partition(nums, i, p-1, k)
	}
	return partition(nums, p+1, j, k)
}
