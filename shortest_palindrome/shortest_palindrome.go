package shortest_palindrome

func shortestPalindrome(s string) string {
	matrix := make([][]bool, len(s))
	for i := range matrix {
		matrix[i] = make([]bool, len(s))
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := i; j < len(matrix[i]); j++ {
			if j == i || s[i] == s[j] && (i+1 == j || matrix[i+1][j-1]) {
				matrix[i][j] = true
			}
		}
	}

	p := len(s) - 1
	for ; p >= 0; p-- {
		if matrix[0][p] {
			break
		}
	}
	return reverse(s[p+1:]) + s
}

func reverse(s string) string {
	bytes := make([]byte, len(s))
	for i := range s {
		bytes[len(s)-i-1] = s[i]
	}
	return string(bytes)
}
