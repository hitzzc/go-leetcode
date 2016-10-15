package excel_sheet_column_title

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}
	bytes := []byte{}
	var remainder int
	for n > 0 {
		remainder = (n - 1) % 26
		bytes = append(bytes, 'A'+byte(remainder))
		n = (n - 1) / 26
	}
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
