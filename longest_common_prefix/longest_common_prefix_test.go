package longest_common_prefix

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := [][]string{
		[]string{"", ""},
		[]string{"abc", "abd", "abcde"},
		[]string{"", "abc"},
	}
	results := []string{
		"",
		"ab",
		"",
	}
	caseNums := 3
	for i := 0; i < caseNums; i++ {
		if ret := longestCommonPrefix(tests[i]); ret != results[i] {
			t.Fatalf("case %d failed\nactual: %+v, expect: %+v\n", i, ret, results[i])
		}
	}
}
