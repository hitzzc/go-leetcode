package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	list := []*ListNode{}
	var ptr *ListNode
	for ptr = head; ptr.Next != nil; ptr = ptr.Next {
		list = append(list, ptr)
	}
	list = append(list, ptr)
	index := len(list) - n
	if index == 0 {
		return head.Next
	}
	list[index-1].Next = list[index].Next
	return head
}
