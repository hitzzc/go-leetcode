package swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	newHead.Next, head.Next = head, swapPairs(newHead.Next)
	return newHead
}
