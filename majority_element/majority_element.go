package majority_element

func majorityElement(nums []int) int {
	var elem, count int
	for i := range nums {
		if count == 0 {
			count++
			elem = nums[i]
		} else {
			if nums[i] != elem {
				count--
			} else {
				count++
			}
		}
	}
	return elem
}
