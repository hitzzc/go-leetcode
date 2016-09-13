package zigzag_conversion

const (
	up   string = "up"
	down        = "down"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	runes := []rune(s)
	rows := make([][]rune, numRows)
	i := 0
	status := down
	for index := range runes {
		rows[i] = append(rows[i], runes[index])
		if status == down {
			i++
			if i == numRows-1 {
				status = up
			}
		} else {
			i--
			if i == 0 {
				status = down
			}
		}
	}
	retRunes := []rune{}
	for index := range rows {
		retRunes = append(retRunes, rows[index]...)
	}
	return string(retRunes)
}
