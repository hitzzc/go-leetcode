package search_for_a_range

import (
	"testing"
)

func TestSearchRange(t *testing.T) {
	ret := searchRange([]int{1, 2, 3, 3, 3, 3, 4, 5, 9}, 3)
	println("ret", ret[0], ret[1])
}
