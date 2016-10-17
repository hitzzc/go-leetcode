package majority_element_II

func majorityElement(nums []int) []int {
	var m, n, cm, cn int
	for i := range nums {
		if nums[i] == m {
			cm++
		} else if nums[i] == n {
			cn++
		} else if cm == 0 {
			m = nums[i]
			cm++
		} else if cn == 0 {
			n = nums[i]
			cn++
		} else {
			cm--
			cn--
		}
	}
	cm, cn = 0, 0
	for i := range nums {
		if nums[i] == m {
			cm++
		} else if nums[i] == n {
			cn++
		}
	}
	var ret []int
	if cm > len(nums)/3 {
		ret = append(ret, m)
	}
	if cn > len(nums)/3 {
		ret = append(ret, n)
	}
	return ret
}
