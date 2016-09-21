package longest_valid_parentheses

func longestValidParentheses(s string) int {
	bytes := []byte(s)
	if len(bytes) < 2 {
		return 0
	}
	lengthList := make([]int, len(bytes))
	var max int
	for i := 1; i < len(bytes); i++ {
		if bytes[i] == ')' {
			j := i - lengthList[i-1] - 1
			if j >= 0 && bytes[j] == '(' {
				lengthList[i] = lengthList[i-1] + 2
				if j-1 >= 0 {
					lengthList[i] += lengthList[j-1]
				}
			}
		}
		if lengthList[i] > max {
			max = lengthList[i]
		}
	}
	return max
}
