package permutation_sequence

func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}
	data := make([]int, n)
	data[0] = 1
	iToS := []byte("123456789")
	for i := 1; i < n; i++ {
		data[i] = data[i-1] * i
	}
	ret := ""
	k--
	for i := n - 1; i >= 0; i-- {
		ret += string(iToS[k/data[i]])
		iToS = append(iToS[:k/data[i]], iToS[k/data[i]+1:]...)
		k = k % data[i]
	}
	return ret
}
