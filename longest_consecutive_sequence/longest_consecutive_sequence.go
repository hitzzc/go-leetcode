package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	visited := map[int]struct{}{}
	existed := map[int]struct{}{}
	for i := range nums {
		existed[nums[i]] = struct{}{}
	}
	max := 0
	for i := range nums {
		count := 0
		left := nums[i]
		for {
			if _, ok := visited[left]; ok {
				break
			} else if _, ok := existed[left]; ok {
				visited[left] = struct{}{}
				count++
				left--
			} else {
				break
			}
		}
		right := nums[i] + 1
		for {
			if _, ok := visited[right]; ok {
				break
			} else if _, ok := existed[right]; ok {
				visited[right] = struct{}{}
				count++
				right++
			} else {
				break
			}
		}
		if count > max {
			max = count
		}
	}
	return max
}
