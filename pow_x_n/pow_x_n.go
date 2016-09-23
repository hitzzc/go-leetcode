package pow_x_n

func myPow(x float64, n int) float64 {
	ret := 1.0
	sign := true
	if n < 0 {
		n = -n
		sign = false
	}
	for n != 0 {
		if n%2 == 1 {
			ret *= x
		}
		n >>= 1
		x *= x
	}

	if !sign {
		return 1 / ret
	}
	return ret
}
