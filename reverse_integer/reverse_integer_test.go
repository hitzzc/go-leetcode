package reverse_integer

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []int{
		123,
		-123,
		0,
		-1,
		1534236469,
	}
	results := []int{
		321,
		-321,
		0,
		-1,
		0,
	}
	caseNum := 5
	for i := 0; i < caseNum; i++ {
		if ret := reverse(tests[i]); ret != results[i] {
			t.Fatalf("case %d failed\nactual: %d, expect: %d\n", i, ret, results[i])
		}
	}
}
