package combination_sum

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return DFS(candidates, 0, target, []int{}, [][]int{})
}

func DFS(candidates []int, start int, target int, solution []int, results [][]int) (rets [][]int) {
	rets = results
	if target < 0 {
		return
	}
	if target == 0 {
		rets = append(rets, solution)
		return
	}
	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		candidate := candidates[i]
		// here we should copy solution to repliaSolution. Directly using slice(solution) in iteration will make results fail. Fucking golang's slice
		repliaSolution := make([]int, len(solution))
		copy(repliaSolution, solution)
		repliaSolution = append(repliaSolution, candidate)
		rets = DFS(candidates, i, target-candidate, repliaSolution, rets)
		repliaSolution = repliaSolution[:len(repliaSolution)-1]
	}
	return
}
