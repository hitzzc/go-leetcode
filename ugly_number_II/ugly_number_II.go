package ugly_number_II

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	records := []int{1}
	i, j, k := 0, 0, 0
	for len(records) < n {
		next := min(2*records[i], min(3*records[j], 5*records[k]))
		if next == 2*records[i] {
			i++
		}
		if next == 3*records[j] {
			j++
		}
		if next == 5*records[k] {
			k++
		}
		records = append(records, next)
	}
	return records[len(records)-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
