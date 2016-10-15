package house_robber

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	array := make([]int, len(nums))
	array[0] = nums[0]
	if array[0] > nums[1] {
		array[1] = array[0]
	} else {
		array[1] = nums[1]
	}
	for i := 2; i < len(nums); i++ {
		if array[i-1] > array[i-2]+nums[i] {
			array[i] = array[i-1]
		} else {
			array[i] = array[i-2] + nums[i]
		}
	}
	return array[len(array)-1]
}
