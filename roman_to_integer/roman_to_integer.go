package roman_to_integer

func romanToInt(s string) int {
	m := map[rune]int{
		rune('I'): 1,
		rune('V'): 5,
		rune('X'): 10,
		rune('L'): 50,
		rune('C'): 100,
		rune('D'): 500,
		rune('M'): 1000,
	}
	runes := []rune(s)
	if len(runes) == 1 {
		return m[runes[0]]
	}
	ret := m[runes[0]]
	for i := 1; i < len(runes); i++ {
		if m[runes[i-1]] < m[runes[i]] {
			ret += m[runes[i]] - 2*m[runes[i-1]]
		} else {
			ret += m[runes[i]]
		}
	}
	return ret
}
