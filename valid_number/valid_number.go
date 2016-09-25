package valid_number

func isNumber(s string) bool {
	bytes := []byte(s)
	point, hasE := false, false
	i, n := 0, len(bytes)
	for ; i < n; i++ {
		if bytes[i] != ' ' {
			break
		}
	}
	if i == n {
		return false
	}
	if bytes[i] == '+' || bytes[i] == '-' {
		i++
	}
	if i == n {
		return false
	}
	head := i
	for ; i < n; i++ {
		switch bytes[i] {
		case '.':
			if point || hasE || i == head && (i == n-1 || !isDigit(bytes[i+1])) {
				return false
			}
			point = true
			continue
		case 'e':
			if hasE || i == head || i == n-1 {
				return false
			}
			i++
			if bytes[i] == '-' || bytes[i] == '+' {
				i++
			}
			if i == n || !isDigit(bytes[i]) {
				return false
			}
			hasE = true
			continue
		}
		if bytes[i] == ' ' {
			for ; i < n; i++ {
				if bytes[i] != ' ' {
					return false
				}
			}
			return true
		}
		if !isDigit(bytes[i]) {
			return false
		}
	}
	return true
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}
