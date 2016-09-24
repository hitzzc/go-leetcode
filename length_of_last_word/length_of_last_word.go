package length_of_last_word

func lengthOfLastWord(s string) int {
	bytes := []byte(s)
	lastIndex := len(bytes) - 1
	for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] != ' ' {
			lastIndex = i
			break
		}
	}
	bytes = bytes[:lastIndex+1]
	for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] == ' ' {
			return len(bytes) - 1 - i
		}
	}
	return len(bytes)
}
