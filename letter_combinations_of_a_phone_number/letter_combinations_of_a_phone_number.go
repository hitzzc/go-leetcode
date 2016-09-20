package letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
		"0": []string{" "},
	}
	slices := [][]string{}
	for index := range digits {
		digit := string(digits[index])
		if _, ok := m[digit]; ok {
			slices = append(slices, m[digit])
		}
	}
	return generateCombinations(slices)
}

func generateCombinations(slices [][]string) (ret []string) {
	if len(slices) == 0 {
		return
	}
	if len(slices) == 1 {
		return slices[0]
	}
	for _, letter := range slices[0] {
		tmpRets := generateCombinations(slices[1:])
		for _, tmp := range tmpRets {
			ret = append(ret, letter+tmp)
		}
	}
	return
}
