package string_to_integer

// just consider 32-bit, fucking leetcode.com's test cases
const (
	MAX int32 = 1<<31 - 1
	MIN int32 = -1*MAX - 1
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	runes := []rune(str)
	var ret int32
	var neg bool
	i := 0
	for ; runes[i] == rune(' '); i++ {
	}
	if runes[i] == rune('-') || runes[i] == rune('+') {
		neg = runes[i] == rune('-')
		i++
	}

	for ; i < len(runes); i++ {
		if val, ok := toInt(runes[i]); !ok {
			break
		} else {
			if neg {
				if -ret < (MIN+val)/10 {
					return int(MIN)
				}
			} else {
				if ret > (MAX-val)/10 {
					return int(MAX)
				}
			}
			ret = 10*ret + val
		}
	}
	if neg {
		return int(-ret)
	}
	return int(ret)
}

func toInt(r rune) (int32, bool) {
	ret := int32(r - rune('0'))
	if ret >= 0 && ret <= 9 {
		return ret, true
	}
	return 0, false
}
