package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var tail, next *ListNode
	for head != nil && head.Next != nil {
		next = head.Next
		head.Next = tail
		tail = head
		head = next
	}
	if head != nil {
		head.Next = tail
	}
	return head
}
