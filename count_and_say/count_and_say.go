package count_and_say

func countAndSay(n int) string {
	if n <= 0 {
		return ""
	} else if n == 1 {
		return "1"
	}
	rets := countAndSayInts(n)
	bytes := make([]byte, len(rets))
	intTobyte := []byte("123456789")
	for index, ret := range rets {
		bytes[index] = intTobyte[ret-1]
	}
	return string(bytes)
}

func countAndSayInts(n int) (rets []int) {
	if n < 1 {
		return []int{}
	} else if n == 1 {
		return []int{1}
	}
	rets = []int{1}
	for i := 2; i <= n; i++ {
		tmps := []int{}
		pre, cnt := -1, 0
		for index, val := range rets {
			if index == 0 {
				pre = val
				cnt = 1
				continue
			}
			if val != pre {
				tmps = append(tmps, cnt, pre)
				pre = val
				cnt = 1
			} else {
				cnt++
			}
		}
		tmps = append(tmps, cnt, pre)
		rets = tmps
	}
	return
}
