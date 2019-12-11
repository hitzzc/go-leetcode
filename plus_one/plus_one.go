package plus_one

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits)-1; i >= 0; i-- {
		sum := carry + digits[i]
		if sum < 10 {
			digits[i] = sum
			return digits
		}
		digits[i] = sum % 10
		carry = 1
	}
	if carry == 1 {
		return append([]int{carry}, digits...)
	}
	return digits
}
