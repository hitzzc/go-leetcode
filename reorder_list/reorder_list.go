package reorder_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	var tail *ListNode
	for fast != nil && fast.Next != nil {
		tail = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if tail != nil {
		tail.Next = nil
	}
	var secondHead *ListNode
	fast = slow
	for fast != nil {
		slow = fast
		fast = fast.Next
		slow.Next = secondHead
		secondHead = slow
	}
	firstHead := head
	var p1, p2 *ListNode
	for firstHead != nil && secondHead != nil {
		p1 = firstHead.Next
		p2 = secondHead.Next
		firstHead.Next = secondHead
		secondHead.Next = p1
		if p1 == nil {
			secondHead.Next = p2
			break
		}
		firstHead = p1
		secondHead = p2
	}
	return
}
