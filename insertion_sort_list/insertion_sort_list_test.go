package insertion_sort_list

import (
	"testing"
)

func (this *ListNode) insert(val int) {
	p := this
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &ListNode{Val: val}
}

func (this *ListNode) print() {
	p := this
	for p != nil {
		println(p.Val)
		p = p.Next
	}
}

func TestInsertionSortList(t *testing.T) {
	head := &ListNode{Val: 3}
	head.insert(2)
	head.insert(1)
	head.insert(4)
	head.insert(5)
	head.insert(1)
	head = insertionSortList(head)
	head.print()
}
