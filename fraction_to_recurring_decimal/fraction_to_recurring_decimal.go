package fraction_to_recurring_decimal

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	m := []byte("0123456789")
	var ret string
	var neg bool
	if numerator < 0 {
		numerator *= -1
		neg = !neg
	}
	if denominator < 0 {
		denominator *= -1
		neg = !neg
	}
	if neg {
		ret = "-"
	}
	ret += strconv.Itoa(numerator / denominator)
	remainder := numerator % denominator
	if remainder == 0 {
		return ret
	}
	ret += "."
	floatBytes := []byte{}
	indexMap := map[int]int{}
	for remainder != 0 {
		if idx, ok := indexMap[remainder]; ok {
			return ret + string(floatBytes[:idx]) + "(" + string(floatBytes[idx:]) + ")"
		}
		indexMap[remainder] = len(floatBytes)
		remainder *= 10
		v := m[remainder/denominator]
		remainder %= denominator
		floatBytes = append(floatBytes, v)
	}
	return ret + string(floatBytes)
}
