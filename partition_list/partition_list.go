package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	lower, bigger := &ListNode{}, &ListNode{}
	ptr1, ptr2 := lower, bigger
	for head != nil {
		if head.Val < x {
			ptr1.Next = head
			ptr1 = ptr1.Next
		} else {
			ptr2.Next = head
			ptr2 = ptr2.Next
		}
		head = head.Next
	}
	ptr2.Next = nil
	ptr1.Next = bigger.Next
	return lower.Next
}
