package combination_sum_II

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
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
		repliaSolution := make([]int, len(solution))
		copy(repliaSolution, solution)
		repliaSolution = append(repliaSolution, candidate)
		rets = DFS(candidates, i+1, target-candidate, repliaSolution, rets)
		repliaSolution = repliaSolution[:len(repliaSolution)-1]
	}
	return
}
