package subsets_II

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	solution := []int{}
	ret := [][]int{[]int{}}
	ret = append(ret, subsetsI(nums, solution, -1)...)
	return ret
}

func subsetsI(nums []int, solution []int, idx int) [][]int {
	ret := [][]int{}
	for i := idx + 1; i < len(nums); i++ {
		if i > idx+1 && nums[i] == nums[i-1] {
			continue
		}
		solution = append(solution, nums[i])
		cp1 := make([]int, len(solution))
		copy(cp1, solution)
		ret = append(ret, cp1)
		tmps := subsetsI(nums, solution, i)
		ret = append(ret, tmps...)
		solution = solution[:len(solution)-1]
	}
	return ret
}
