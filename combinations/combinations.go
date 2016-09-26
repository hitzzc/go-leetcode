package combinations

func combine(n int, k int) [][]int {
	rets := [][]int{}
	if n <= 0 || k <= 0 || n < k {
		return rets
	}
	currentInt := make([]int, k)
	solution := []int{}
	depth := -1
	maxDepth := k - 1
	for {
	L:
		for depth < maxDepth {
			for {
				childDepth := depth + 1
				preInt := currentInt[childDepth]
				if preInt == n {
					currentInt[childDepth] = 0
					break L
				}
				currentInt[childDepth]++
				if depth == -1 || currentInt[childDepth] > currentInt[depth] {
					solution = append(solution, currentInt[childDepth])
					break
				}
			}
			depth++
		}
		if depth == maxDepth {
			cp := make([]int, len(solution))
			copy(cp, solution)
			rets = append(rets, cp)
		}
		if depth == -1 {
			break
		}
		depth--
		solution = solution[:len(solution)-1]
	}
	return rets
}
