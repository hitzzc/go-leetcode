package course_schedule_II

func findOrder(numCourses int, prerequisites [][]int) []int {
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
	ret := []int{}
	for len(queue) != 0 {
		head := queue[0]
		queue = queue[1:]
		for i := range graph[head] {
			inDegree[graph[head][i]]--
			if inDegree[graph[head][i]] == 0 {
				queue = append(queue, graph[head][i])
			}
		}
		ret = append(ret, head)
	}
	if len(ret) != numCourses {
		return []int{}
	}
	return ret
}
