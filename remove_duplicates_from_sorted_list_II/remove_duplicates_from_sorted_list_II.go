package remove_duplicates_from_sorted_list_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	helper := &ListNode{Next: head}
	pre, current := helper, helper.Next
	for current != nil {
		for current.Next != nil && current.Next.Val == current.Val {
			current = current.Next
		}
		if pre.Next == current {
			pre = pre.Next
		} else {
			pre.Next = current.Next
		}
		current = current.Next
	}
	return helper.Next
}
