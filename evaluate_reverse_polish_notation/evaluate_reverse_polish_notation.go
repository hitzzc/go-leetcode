package evaluate_reverse_polish_notation

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	cache := []int{}
	var num1, num2, num int
	for i := range tokens {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			num2 = cache[len(cache)-1]
			num1 = cache[len(cache)-2]
			cache[len(cache)-2] = calculate(num1, num2, tokens[i])
			cache = cache[:len(cache)-1]
		} else {
			num, _ = strconv.Atoi(tokens[i])
			cache = append(cache, num)
		}
	}
	return cache[0]
}

func calculate(num1, num2 int, op string) int {
	if op == "+" {
		return num1 + num2
	}
	if op == "-" {
		return num1 - num2
	}
	if op == "*" {
		return num1 * num2
	}
	if op == "/" {
		return num1 / num2
	}
	return 0
}
