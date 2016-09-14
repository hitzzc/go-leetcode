package reverse_integer

func reverse(x int) int {
	sign := 1
	tmp := x
	if x < 0 {
		sign = -1
		tmp = -1 * x
	}
	list := []int{}
	for tmp > 9 {
		list = append(list, tmp%10)
		tmp = tmp / 10
	}
	list = append(list, tmp)
	ret := 0
	for i := 0; i < len(list); i++ {
		ret = 10*ret + list[i]
	}
	return ret * sign
}
