package number_of_islands

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var ret int
	visited := map[int]bool{}
	var k int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}
			k = hash(i, j, len(grid[0]))
			if _, ok := visited[k]; ok {
				continue
			}
			ret++
			queue := []int{k}
			for len(queue) != 0 {
				top := queue[0]
				queue = queue[1:]
				row, col := top/len(grid[0]), top%len(grid[0])
				if row > 0 && grid[row-1][col] == '1' {
					k = hash(row-1, col, len(grid[0]))
					if _, ok := visited[k]; !ok {
						queue = append(queue, k)
						visited[k] = true
					}
				}
				if row < len(grid)-1 && grid[row+1][col] == '1' {
					k = hash(row+1, col, len(grid[0]))
					if _, ok := visited[k]; !ok {
						queue = append(queue, k)
						visited[k] = true
					}
				}
				if col > 0 && grid[row][col-1] == '1' {
					k = hash(row, col-1, len(grid[0]))
					if _, ok := visited[k]; !ok {
						queue = append(queue, k)
						visited[k] = true
					}
				}
				if col < len(grid[0])-1 && grid[row][col+1] == '1' {
					k = hash(row, col+1, len(grid[0]))
					if _, ok := visited[k]; !ok {
						queue = append(queue, k)
						visited[k] = true
					}
				}
			}
		}
	}
	return ret
}

func hash(row, col, length int) int {
	return row*length + col
}
