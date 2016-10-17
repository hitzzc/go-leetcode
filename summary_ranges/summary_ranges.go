package summary_ranges

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	var ret []string
	if len(nums) == 0 {
		return ret
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	var i int
	var j int = 1
	for ; j < len(nums); j++ {
		if nums[j] == nums[j-1]+1 {
			continue
		}
		if nums[j] != nums[j-1]+1 {
			if j-1 == i {
				ret = append(ret, strconv.Itoa(nums[i]))
			} else {
				ret = append(ret, strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[j-1]))
			}
			i = j
		}
	}
	if j-1 == i {
		ret = append(ret, strconv.Itoa(nums[i]))
	} else {
		ret = append(ret, strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[j-1]))
	}
	return ret
}
