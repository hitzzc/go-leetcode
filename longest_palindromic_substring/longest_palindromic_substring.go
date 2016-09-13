package longest_palindromic_substring

func longestPalindrome(s string) string {
	runes := []rune(s)
	length := len(runes)
	if length <= 1 {
		return s
	}
	maxRune := runes[0:1]
	for i := 0; i < length-1; i++ {
		tmpRune := findPalindrome(runes, i, i)
		if len(tmpRune) > len(maxRune) {
			maxRune = tmpRune
		}
		if runes[i] == runes[i+1] {
			tmpRune = findPalindrome(runes, i, i+1)
			if len(tmpRune) > len(maxRune) {
				maxRune = tmpRune
			}
		}
	}
	return string(maxRune)
}

func findPalindrome(subRune []rune, i, j int) []rune {
	m, n := i-1, j+1
	for ; m >= 0 && n <= len(subRune)-1 && subRune[m] == subRune[n]; m, n = m-1, n+1 {
	}
	return subRune[m+1 : n]
}
