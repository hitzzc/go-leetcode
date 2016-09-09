package add_two_numbers

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	left := [][]int{
		[]int{3, 2, 4},
		[]int{8, 7, 3, 1},
		[]int{8, 1},
	}
	right := [][]int{
		[]int{5, 8, 5},
		[]int{2, 1},
		[]int{1, 2},
	}
	results := [][]int{
		[]int{8, 0, 0, 1},
		[]int{0, 9, 3, 1},
		[]int{9, 3},
	}
	caseNum := 3
	for i := 0; i < caseNum; i++ {
		l, r := generateNodeList(left[i]), generateNodeList(right[i])
		if ret := addTwoNumbers(l, r); !equal(ret, generateNodeList(results[i])) {
			t.Fatalf("case %d fails: %v\n", i, ret.Val)
		}
	}
}

func generateNodeList(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	head := new(ListNode)
	walk := head
	for index, val := range list {
		walk.Val = val
		if index != len(list)-1 {
			walk.Next = new(ListNode)
			walk = walk.Next
		}
	}
	return head
}

func equal(left, right *ListNode) bool {
	tmp1, tmp2 := left, right
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val != tmp2.Val {
			return false
		}
		tmp1, tmp2 = tmp1.Next, tmp2.Next
	}
	if tmp1 != nil || tmp2 != nil {
		return false
	}
	return true
}
