package insertion_sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fake := &ListNode{Next: head}
	tail := head
	var cur *ListNode
	for tail.Next != nil {
		cur = tail.Next
		p := fake
		for ; p.Next != cur; p = p.Next {
			if p.Next.Val >= cur.Val {
				tmp := p.Next
				p.Next = cur
				tail.Next = cur.Next
				cur.Next = tmp
				break
			}
		}
		if p == tail {
			tail = cur
		}
	}
	return fake.Next
}
