package maximum_gap

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	max, min := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	bucketSize := (max-min)/len(nums) + 1
	bucket := make([][]int, len(nums))
	for i := range nums {
		idx := (nums[i] - min) / bucketSize
		if len(bucket[idx]) == 0 {
			bucket[idx] = []int{nums[i], nums[i]}
		} else {
			if nums[i] < bucket[idx][0] {
				bucket[idx][0] = nums[i]
			}
			if nums[i] > bucket[idx][1] {
				bucket[idx][1] = nums[i]
			}
		}
	}
	ret := bucket[0][1] - bucket[0][0]
	pre := 0
	for i := 1; i < len(nums); i++ {
		if len(bucket[i]) == 0 {
			continue
		}
		if bucket[i][0]-bucket[pre][1] > ret {
			ret = bucket[i][0] - bucket[pre][1]
		}
		pre = i
	}
	return ret
}
