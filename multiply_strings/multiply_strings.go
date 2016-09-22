package multiply_strings

func multiply(num1 string, num2 string) string {
	num1Bytes, num2Bytes := []byte(num1), []byte(num2)
	num1Ints, num2Ints := BytesToInts(num1Bytes), BytesToInts(num2Bytes)
	var ret []int
	for i := 0; i < len(num2Ints); i++ {
		tmp := oneMultiply(num1Ints, num2Ints[i])
		ret = intsPlus(append(ret, 0), tmp)
	}
	if len(ret) == 0 {
		return "0"
	}
	var i int
	for i = range ret {
		if ret[i] != 0 {
			break
		}
	}
	ret = ret[i:]
	return string(IntsToBytes(ret))
}

func BytesToInts(bytes []byte) (ints []int) {
	ints = make([]int, len(bytes))
	for i, b := range bytes {
		ints[i] = int(b - '0')
	}
	return
}

func IntsToBytes(ints []int) (bytes []byte) {
	if len(ints) == 0 {
		return []byte{'0'}
	}
	m := []byte("0123456789")
	bytes = make([]byte, len(ints))
	for i, val := range ints {
		bytes[i] = m[val]
	}
	return
}

func oneMultiply(num1 []int, num2 int) []int {
	carry := 0
	ret := make([]int, len(num1)+1)
	for i := len(num1) - 1; i >= 0; i-- {
		dot := num1[i]*num2 + carry
		if dot > 9 {
			carry = dot / 10
			dot = dot % 10
		} else {
			carry = 0
		}
		ret[i+1] = dot
	}
	if carry == 0 {
		return ret[1:]
	}
	ret[0] = carry
	return ret
}

func intsPlus(num1, num2 []int) []int {
	if len(num1) == 0 {
		return num2
	}
	if len(num2) == 0 {
		return num1
	}
	var short, long []int
	if len(num1) < len(num2) {
		short = num1
		long = num2
	} else {
		short = num2
		long = num1
	}
	newShort := make([]int, len(long)-len(short))
	newShort = append(newShort, short...)
	ret := make([]int, len(long)+1)
	carry := 0
	for i := len(newShort) - 1; i >= 0; i-- {
		sum := carry + newShort[i] + long[i]
		if sum > 9 {
			carry = sum / 10
			sum = sum % 10
		} else {
			carry = 0
		}
		ret[i+1] = sum
	}
	if carry == 0 {
		return ret[1:]
	}
	ret[0] = carry
	return ret
}
