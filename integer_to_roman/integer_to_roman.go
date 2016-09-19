package integer_to_roman

func intToRoman(num int) string {
	romanMap := [][]string{
		[]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		[]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		[]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		[]string{"", "M", "MM", "MMM"},
	}
	digits, tmp := 0, num
	var ret string
	for tmp > 0 {
		remain := tmp % 10
		ret = romanMap[digits][remain] + ret
		tmp = tmp / 10
		digits++
	}
	return ret
}
