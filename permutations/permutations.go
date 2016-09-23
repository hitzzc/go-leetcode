package permutations

import (
//"fmt"
)

// recursion
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

// no recursion
func permute1(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	visited := map[int]bool{}
	// depth==-1 means the root
	depth, maxDepth := -1, len(nums)-1
	currentIndex := make([]int, maxDepth+1)
	for i := 0; i < maxDepth; i++ {
		currentIndex[i] = -1
	}
	stack := []int{}
	ret := [][]int{}
	for {
	L:
		for depth < maxDepth {
			for {
				childDepth := depth + 1
				if currentIndex[childDepth] == len(nums)-1 {
					currentIndex[childDepth] = -1
					break L
				}
				currentIndex[childDepth]++
				val := nums[currentIndex[childDepth]]
				if !visited[val] {
					stack = append(stack, val)
					visited[val] = true
					break
				}
			}
			depth++
		}
		if depth == maxDepth {
			tmp := make([]int, len(stack))
			copy(tmp, stack)
			ret = append(ret, tmp)
		}
		depth--
		if depth == -2 {
			break
		}
		visited[stack[len(stack)-1]] = false
		stack = stack[:len(stack)-1]
	}
	return ret
}
