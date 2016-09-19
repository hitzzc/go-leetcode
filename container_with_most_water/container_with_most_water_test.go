package container_with_most_water

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := [][]int{
		[]int{1, 1},
		[]int{1, 2},
		[]int{1, 2, 4, 3},
	}
	results := []int{
		1,
		1,
		4,
	}
	caseNums := 3
	for i := 0; i < caseNums; i++ {
		if ret := maxArea(tests[i]); ret != results[i] {
			t.Fatalf("case %d failed\nactual: %+v, expect: %+v\n", i, ret, results[i])
		}
	}
}
