package largest_number

import (
	"strconv"
)

func largestNumber(nums []int) string {
	stringList := make([]string, len(nums))
	for i := range nums {
		stringList[i] = strconv.Itoa(nums[i])
	}
	quickSort(stringList, 0, len(stringList)-1)
	var ret string
	for i := range stringList {
		ret += stringList[i]
	}
	if len(ret) != 0 && ret[0] == '0' {
		return "0"
	}
	return ret
}

func quickSort(stringList []string, i, j int) {
	if i >= j {
		return
	}
	mid := partition(stringList, i, j)
	quickSort(stringList, i, mid-1)
	quickSort(stringList, mid+1, j)
	return
}

func partition(stringList []string, i, j int) int {
	k := i
	for ; i < j; i++ {
		if less(stringList[i], stringList[j]) {
			continue
		}
		stringList[k], stringList[i] = stringList[i], stringList[k]
		k++
	}
	stringList[k], stringList[j] = stringList[j], stringList[k]
	return k
}

func less(s1, s2 string) bool {
	return s1+s2 <= s2+s1
}
