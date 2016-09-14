package reverse_integer

// fucking leetcode's int32 overflow test cases
const (
	MAX int32 = 1<<31 - 1
	MIN       = -MAX - 1
)

func reverse(x int) int {
	var sign int32 = 1
	tmp := x
	if x < 0 {
		sign = -1
		tmp = -1 * x
	}
	list := []int32{}
	for tmp > 9 {
		list = append(list, int32(tmp%10))
		tmp = tmp / 10
	}
	list = append(list, int32(tmp))
	var ret int32
	for i := 0; i < len(list); i++ {
		if sign == 1 {
			if ret > (MAX-list[i])/10 {
				return 0
			}
		} else {
			if -ret < (MIN+list[i])/10 {
				return 0
			}
		}
		ret = 10*ret + list[i]
	}
	return int(ret * sign)
}
