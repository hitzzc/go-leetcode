package power_of_two

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	for n > 1 {
		if n&1 != 0 {
			return false
		}
		n >>= 1
	}
	return true
}
