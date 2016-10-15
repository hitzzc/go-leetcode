package factorial_trailing_zeroes

func trailingZeroes(n int) int {
	ret := 0
	for n > 0 {
		ret += n / 5
		n /= 5
	}
	return ret
}
