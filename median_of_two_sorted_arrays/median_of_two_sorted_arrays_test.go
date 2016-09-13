package median_of_two_sorted_arrays

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := [][]int{
		[]int{1, 3},
		[]int{1, 2, 3},
		[]int{4},
	}
	nums2 := [][]int{
		[]int{2},
		[]int{4, 5, 6},
		[]int{1, 2, 3, 7, 8, 9, 10, 11},
	}
	results := []float64{
		2,
		3.5,
		7,
	}
	caseNums := 3
	for i := 0; i < caseNums; i++ {
		if ret := findMedianSortedArrays(nums1[i], nums2[i]); ret != results[i] {
			t.Fatalf("case %d failed.\nexpect: %f, actual: %f\n", i, results[i], ret)
		}
	}
}
