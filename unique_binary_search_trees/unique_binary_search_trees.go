package unique_binary_search_trees

func numTrees(n int) int {
	if n <= 1 {
		return n
	}
	array := make([]int, n+1)
	array[0], array[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			array[i] += array[j] * array[i-j-1]
		}
	}
	return array[n]
}
