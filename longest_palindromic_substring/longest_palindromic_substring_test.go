package longest_palindromic_substring

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []string{
		"a",
		"aaaaa",
		"abababab",
		"abcfcdfgfdc",
	}
	results := []string{
		"a",
		"aaaaa",
		"abababa",
		"cdfgfdc",
	}
	caseNums := 4
	for i := 0; i < caseNums; i++ {
		if ret := longestPalindrome(tests[i]); ret != results[i] {
			t.Fatalf("case %d failed\nactual: %s, expect: %s\n", i, ret, results[i])
		}
	}
}
