package n_queens_II

func totalNQueens(n int) int {
	ret, depth, maxDepth := 0, -1, n-1
	nextIndex := make([]int, n)
	stack := []int{}
	for {
	L:
		for depth < maxDepth {
			for {
				childDepth := depth + 1
				if nextIndex[childDepth] == n {
					nextIndex[childDepth] = 0
					break L
				}
				index := nextIndex[childDepth]
				nextIndex[childDepth]++
				if isValid(stack, childDepth, index) {
					stack = append(stack, index)
					break
				}
			}
			depth++
		}
		if depth == maxDepth {
			ret++
		}
		if depth == -1 {
			break
		}
		depth--
		stack = stack[:len(stack)-1]
	}
	return ret
}

func isValid(stack []int, index, val int) bool {
	for i, v := range stack {
		if v == val || i-index == v-val || i-index == val-v {
			return false
		}
	}
	return true
}
