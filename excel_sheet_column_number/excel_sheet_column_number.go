package excel_sheet_column_number

func titleToNumber(s string) int {
	bytes := []byte(s)
	var ret int
	for i := range bytes {
		ret = 26*ret + int(bytes[i]-'A'+1)
	}
	return ret
}
