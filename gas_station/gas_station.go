package gas_station

func canCompleteCircuit(gas []int, cost []int) int {
	var total, sum, diff, start int
	for i := range gas {
		diff = gas[i] - cost[i]
		total += diff
		if sum < 0 {
			sum = diff
			start = i
		} else {
			sum += diff
		}
	}
	if total < 0 {
		return -1
	} else {
		return start
	}
}
