package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre, current := head, head
	for current != nil {
		for current.Next != nil && current.Val == current.Next.Val {
			current = current.Next
		}
		pre.Next = current.Next
		pre, current = current.Next, current.Next
	}
	return head
}
