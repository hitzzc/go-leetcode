package subsets

func subsets(nums []int) [][]int {
	rets := [][]int{}
	if len(nums) == 0 {
		return rets
	}
	currentIndex := make([]int, len(nums))
	for i := range currentIndex {
		currentIndex[i] = -1
	}
	solution := []int{}
	depth, maxDepth := -1, len(nums)-1
	for {
		finish := false
	L:
		for depth < maxDepth {
			for {
				childDepth := depth + 1
				preIndex := currentIndex[childDepth]
				if preIndex == len(nums)-1 {
					finish = true
					currentIndex[childDepth] = -1
					break L
				}
				currentIndex[childDepth]++
				if depth == -1 || nums[currentIndex[depth]] < nums[currentIndex[childDepth]] {
					solution = append(solution, nums[currentIndex[childDepth]])
					break
				}
			}
			depth++
		}
		if finish || depth == maxDepth {
			cp := make([]int, len(solution))
			copy(cp, solution)
			rets = append(rets, cp)
			finish = false
		}
		if depth == -1 {
			break
		}
		depth--
		solution = solution[:len(solution)-1]
	}
	return rets
}
