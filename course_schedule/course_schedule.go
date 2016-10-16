package course_schedule

// BFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}
	inDegree := make([]int, numCourses)
	for i := range prerequisites {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
		inDegree[prerequisites[i][0]]++
	}
	queue := []int{}
	for i := range inDegree {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		head := queue[0]
		queue = queue[1:]
		for i := range graph[head] {
			inDegree[graph[head][i]]--
			if inDegree[graph[head][i]] == 0 {
				queue = append(queue, graph[head][i])
			}
		}
	}
	for i := range inDegree {
		if inDegree[i] != 0 {
			return false
		}
	}
	return true
}

// DFS
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}
	for i := range prerequisites {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
	}
	path := make([]bool, numCourses)
	visited := make([]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		if visited[i] {
			continue
		}
		if hasCycle(i, graph, path, visited) {
			return false
		}
	}
	return true
}

func hasCycle(start int, graph map[int][]int, path []bool, visited []bool) bool {
	for i := range graph[start] {
		if visited[graph[start][i]] {
			continue
		}
		if path[graph[start][i]] {
			return true
		}
		path[graph[start][i]] = true
		if hasCycle(graph[start][i], graph, path, visited) {
			return true
		}
		path[graph[start][i]] = false
	}
	visited[start] = true
	return false
}
