package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	bits := 1
	tmp := x
	for {
		if tmp < 10 {
			break
		}
		tmp = tmp / 10
		bits++
	}
	i, j := 1, bits
	for i < j {
		left := (x / divisor(i)) % 10
		right := (x / divisor(j)) % 10
		if left != right {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}

func divisor(bits int) int {
	ret := 1
	for i := 0; i < bits-1; i++ {
		ret = ret * 10
	}
	return ret
}
