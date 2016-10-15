package happy_number

func isHappy(n int) bool {
	happened := map[int]bool{}
	for n != 1 {
		n = squares(n)
		if happened[n] {
			return false
		}
		happened[n] = true
	}
	return true
}

func squares(n int) (ret int) {
	for n > 0 {
		num := n % 10
		ret += num * num
		n /= 10
	}
	return
}
