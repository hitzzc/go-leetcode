package add_binary

func addBinary(a string, b string) string {
	aBytes, bBytes := []byte(a), []byte(b)
	var carry, sum byte = '0', '0'
	short, long := aBytes, bBytes
	if len(aBytes) > len(bBytes) {
		short, long = bBytes, aBytes
	}
	newShort := make([]byte, len(long)-len(short))
	for i := range newShort {
		newShort[i] = '0'
	}
	newShort = append(newShort, short...)
	ret := make([]byte, len(long)+1)
	for i := len(long) - 1; i >= 0; i-- {
		sum, carry = add(newShort[i], long[i], carry)
		ret[i+1] = sum
	}
	if carry == '1' {
		ret[0] = carry
	} else {
		ret = ret[1:]
	}
	return string(ret)
}

func add(a, b, c byte) (byte, byte) {
	if a == '0' && b == '0' {
		return c, '0'
	}
	if a == '1' && b == '1' {
		return c, '1'
	}
	if c == '1' {
		return '0', '1'
	}
	return '1', '0'
}
