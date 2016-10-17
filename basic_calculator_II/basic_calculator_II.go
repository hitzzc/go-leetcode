package basic_calculator_II

func calculate(s string) int {
	numStack := []int{}
	opStack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = 10*num + int(s[i]-'0')
			}
			numStack = append(numStack, num)
			i--
		} else if s[i] == ' ' {
			continue
		} else {
			for len(opStack) != 0 && priority(opStack[len(opStack)-1]) >= priority(s[i]) {
				num := operate(numStack[len(numStack)-2], numStack[len(numStack)-1], opStack[len(opStack)-1])
				numStack[len(numStack)-2] = num
				numStack = numStack[:len(numStack)-1]
				opStack = opStack[:len(opStack)-1]
			}
			opStack = append(opStack, s[i])
		}
	}
	for len(opStack) != 0 {
		num := operate(numStack[len(numStack)-2], numStack[len(numStack)-1], opStack[len(opStack)-1])
		numStack[len(numStack)-2] = num
		numStack = numStack[:len(numStack)-1]
		opStack = opStack[:len(opStack)-1]
	}
	return numStack[0]
}

func priority(op byte) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

func operate(i, j int, op byte) int {
	switch op {
	case '+':
		return i + j
	case '-':
		return i - j
	case '*':
		return i * j
	case '/':
		return i / j
	}
	return 0
}
