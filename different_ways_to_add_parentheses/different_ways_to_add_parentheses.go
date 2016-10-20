package different_ways_to_add_parentheses

import (
	"strconv"
)

func diffWaysToCompute(input string) []int {
	return helper(input, map[string][]int{})
}

func helper(input string, cache map[string][]int) []int {
	if v, ok := cache[input]; ok {
		return v
	}
	ret := []int{}
	for i := range input {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			left := helper(input[:i], cache)
			right := helper(input[i+1:], cache)
			for j := range left {
				for k := range right {
					switch input[i] {
					case '+':
						ret = append(ret, left[j]+right[k])
					case '-':
						ret = append(ret, left[j]-right[k])
					case '*':
						ret = append(ret, left[j]*right[k])
					}
				}
			}
		}
	}
	if len(ret) == 0 {
		num, _ := strconv.Atoi(input)
		ret = append(ret, num)
	}
	cache[input] = ret
	return ret
}
