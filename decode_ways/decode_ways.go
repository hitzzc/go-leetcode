package decode_ways

func numDecodings(s string) int {
	bytes := []byte(s)
	return numDecodingsI(bytes)
}

func numDecodingsI(bytes []byte) (ret int) {
	length := len(bytes)
	if length == 0 {
		return 0
	}
	if length == 1 {
		if bytes[0] == '0' {
			return 0
		}
		return 1
	}
	if bytes[length-1] != '0' {
		ret = numDecodingsI(bytes[:length-1])
		if ret == 0 {
			return 0
		}
	}
	if (bytes[length-2] == '1' && bytes[length-1] >= '0' && bytes[length-1] <= '9') || (bytes[length-2] == '2' && bytes[length-1] >= '0' && bytes[length-1] <= '6') {
		if length == 2 {
			ret++
		} else {
			ret += numDecodingsI(bytes[:length-2])
		}
	} else if bytes[length-2] != '1' && bytes[length-2] != '2' && bytes[length-1] == '0' {
		ret = 0
	}
	return
}
