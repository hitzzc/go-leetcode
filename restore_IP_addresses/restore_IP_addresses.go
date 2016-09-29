package restore_IP_addresses

import (
	"strings"
)

func restoreIpAddresses(s string) []string {
	solution := [][]byte{}
	tmpRets := restoreIpAddressesI([]byte(s), solution, 0)
	ret := make([]string, len(tmpRets))
	for i := range ret {
		tmpStringList := make([]string, len(tmpRets[i]))
		for j := range tmpStringList {
			tmpStringList[j] = string(tmpRets[i][j])
		}
		ret[i] = strings.Join(tmpStringList, ".")
	}
	return ret
}

func restoreIpAddressesI(bytes []byte, solution [][]byte, count int) (ret [][][]byte) {
	if len(bytes) > 3*(4-count) {
		return
	}
	if count == 4 {
		cp := make([][]byte, len(solution))
		for i := range cp {
			cp[i] = make([]byte, len(solution[i]))
			copy(cp[i], solution[i])
		}
		ret = append(ret, cp)
		return
	}
	for i := 1; i <= 3 && i <= len(bytes); i++ {
		if !isVaild(bytes[:i]) {
			break
		}
		solution = append(solution, bytes[:i])
		ret = append(ret, restoreIpAddressesI(bytes[i:], solution, count+1)...)
		solution = solution[:len(solution)-1]
	}
	return ret
}

func isVaild(bytes []byte) bool {
	if len(bytes) == 0 {
		return false
	}
	if len(bytes) == 1 {
		return true
	}
	if len(bytes) == 2 {
		if bytes[0] != '0' {
			return true
		}
		return false
	}
	if len(bytes) == 3 {
		if bytes[0] == '1' || bytes[0] == '2' && (bytes[1] < '5' || bytes[1] == '5' && bytes[2] <= '5') {
			return true
		}
		return false
	}
	return false
}
