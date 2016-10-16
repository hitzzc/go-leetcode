package combinations_sum_III

func combinationSum3(k int, n int) [][]int {
	return DFS(0, 1, k, n, []int{})
}

func DFS(level int, start int, k int, n int, path []int) (ret [][]int) {
	if level == k && n == 0 {
		cp := make([]int, len(path))
		copy(cp, path)
		ret = append(ret, cp)
		return
	}
	for i := start; i <= n && i <= 9; i++ {
		path = append(path, i)
		tmpRet := DFS(level+1, i+1, k, n-i, path)
		ret = append(ret, tmpRet...)
		path = path[:len(path)-1]
	}
	return
}
