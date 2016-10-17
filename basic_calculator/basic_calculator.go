package basic_calculator

func calculate(s string) int {
	ret := 0
	stack := []int{}
	sign := 1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = 10*num + int(s[i]-'0')
			}
			ret += sign * num
			i--
		} else if s[i] == '+' {
			sign = 1
		} else if s[i] == '-' {
			sign = -1
		} else if s[i] == '(' {
			stack = append(stack, ret, sign)
			ret = 0
			sign = 1
		} else if s[i] == ')' {
			signTmp := stack[len(stack)-1]
			retTmp := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			ret = signTmp*ret + retTmp
			sign = 1
		}
	}
	return ret
}
