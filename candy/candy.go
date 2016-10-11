package candy

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	reduceLength, growLength := 0, 1

	pre := 1
	ret := 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] < ratings[i-1] {
			reduceLength++
			if reduceLength >= pre {
				ret++
			}
			ret += reduceLength
			growLength = 1
		} else {
			cur := 1
			if ratings[i] > ratings[i-1] {
				cur = growLength + 1
			}
			growLength = cur
			ret += cur
			reduceLength = 0
			pre = cur
		}
	}
	return ret
}
