package valid_palindrome

func isPalindrome(s string) bool {
	bytes := []byte(s)
	i, j := 0, len(bytes)-1
L:
	for i < j {
		for {
			if i >= j {
				break L
			}
			if !isAlpha(bytes[i]) {
				i++
			} else {
				break
			}
		}
		for {
			if i >= j {
				break L
			}
			if !isAlpha(bytes[j]) {
				j--
			} else {
				break
			}
		}
		if toLower(bytes[i]) != toLower(bytes[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return 'a' + b - 'A'
	}
	return b
}

func isAlpha(b byte) bool {
	if 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || '0' <= b && b <= '9' {
		return true
	}
	return false
}
