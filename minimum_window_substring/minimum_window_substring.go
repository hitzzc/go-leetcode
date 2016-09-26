package minimum_window_substring

func minWindow(s string, t string) string {
	sBytes, tBytes := []byte(s), []byte(t)
	if len(sBytes) == 0 || len(tBytes) == 0 || len(sBytes) < len(tBytes) {
		return ""
	}
	sMap := map[byte]int{}
	tMap := map[byte]int{}
	for _, b := range tBytes {
		tMap[b]++
	}
	begin, tmpBegin := 0, 0
	minSize := -1
	letters := 0
	for i := 0; i < len(sBytes); i++ {
		b := sBytes[i]
		sMap[b]++
		if _, ok := tMap[b]; ok && sMap[b] < tMap[b] {
			letters++
		}
		if letters == len(tBytes) {
			for tmpBegin < i {
				if _, ok := tMap[sBytes[tmpBegin]]; ok {
					if sMap[sBytes[tmpBegin]] == tMap[sBytes[tmpBegin]] {
						break
					}
					sMap[sBytes[tmpBegin]]--
				}
				tmpBegin++
			}
			if minSize == -1 || i-tmpBegin < minSize {
				minSize = i - tmpBegin
				begin = tmpBegin
			}
		}
		letters--
		tmpBegin++
		sMap[sBytes[tmpBegin]]--
	}
	if minSize == -1 {
		return ""
	}
	return s[begin : begin+minSize+1]
}
