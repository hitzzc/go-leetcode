package generate_parentheses

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	var ret []string
	tmps := generateParenthesisI(n-1, n, 1)
	for index := range tmps {
		ret = append(ret, "("+tmps[index])
	}
	return ret
}

func generateParenthesisI(left, right, sum int) []string {
	if left == 0 {
		bytes := make([]byte, right)
		for index := range bytes {
			bytes[index] = ')'
		}
		return []string{string(bytes)}
	}
	var ret []string
	if sum == 0 {
		tmps := generateParenthesisI(left-1, right, 1)
		for index := range tmps {
			ret = append(ret, "("+tmps[index])
		}
	} else {
		tmps := generateParenthesisI(left-1, right, sum+1)
		for index := range tmps {
			ret = append(ret, "("+tmps[index])
		}
		tmps = generateParenthesisI(left, right-1, sum-1)
		for index := range tmps {
			ret = append(ret, ")"+tmps[index])
		}
	}
	return ret
}
