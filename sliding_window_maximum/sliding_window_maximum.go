package sliding_window_maximum

func maxSlidingWindow(nums []int, k int) []int {
	ret := []int{}
	idxArray := []int{}
	for i := range nums {
		for len(idxArray) != 0 && idxArray[0] <= i-k {
			idxArray = idxArray[1:]
		}
		j := len(idxArray) - 1
		for ; j >= 0 && nums[idxArray[j]] <= nums[i]; j-- {
		}
		idxArray = append(idxArray[:j+1], i)
		if i >= k-1 {
			ret = append(ret, nums[idxArray[0]])
		}
	}
	return ret
}
