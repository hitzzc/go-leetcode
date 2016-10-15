package two_sum_II

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	var sum int
	for i < j {
		sum = numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}
