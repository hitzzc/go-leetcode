package four_sum

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				if left > j+1 && nums[left-1] == nums[left] {
					left++
					continue
				}
				if nums[i]+nums[j]+nums[left]+nums[right] == target {
					ret = append(ret, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return ret
}
