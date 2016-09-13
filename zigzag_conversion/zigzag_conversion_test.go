package zigzag_conversion

import (
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []string{
		"A",
		"ABCDE",
		"PAYPALISHIRING",
	}
	rows := []int{
		1,
		3,
		3,
	}
	results := []string{
		"A",
		"AEBDC",
		"PAHNAPLSIIGYIR",
	}
	caseNum := 3
	for i := 0; i < caseNum; i++ {
		if ret := convert(tests[i], rows[i]); ret != results[i] {
			t.Fatalf("case %d failed\nactual: %s, expect: %s\n", i, ret, results[i])
		}
	}
}
