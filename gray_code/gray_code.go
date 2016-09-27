package gray_code

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	tmps := grayCode(n - 1)
	rets := make([]int, 2*len(tmps))
	ones := 1
	for n > 1 {
		ones *= 2
		n--
	}
	for i, tmp := range tmps {
		rets[i] = tmp
		rets[2*len(tmps)-i-1] = tmp + ones
	}
	return rets
}
