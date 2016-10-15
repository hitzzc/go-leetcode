package bitwise_AND_of_numbers_range

func rangeBitwiseAnd(m int, n int) int {
	var cnt uint = 0
	for m != n {
		cnt++
		m >>= 1
		n >>= 1
	}
	return m << cnt
}
