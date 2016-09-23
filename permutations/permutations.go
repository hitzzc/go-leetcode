package permutations

func permute(nums []int) [][]int {
	return traverse(nums, map[int]bool{}, []int{})
}

func traverse(nums []int, m map[int]bool, solution []int) (rets [][]int) {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	if len(solution) == len(nums) {
		replia := make([]int, len(solution))
		copy(replia, solution)
		rets = append(rets, replia)
		return
	}
	for i, num := range nums {
		if m[i] {
			continue
		}
		m[i] = true
		solution = append(solution, num)
		rets = append(rets, traverse(nums, m, solution)...)
		delete(m, i)
		solution = solution[:len(solution)-1]
	}
	return rets
}
