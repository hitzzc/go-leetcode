package reverse_nodes_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 0 {
		return head
	}
	fake := &ListNode{Next: head}
	p := fake
	for p != nil {
		p.Next = reverseKNodes(p.Next, k)
		for i := 0; p != nil && i < k; i++ {
			p = p.Next
		}
	}
	return fake.Next
}

func reverseKNodes(head *ListNode, k int) *ListNode {
	end := head
	for k > 0 && end != nil {
		end = end.Next
		k--
	}
	if k > 0 {
		return head
	}
	ret, pNode := end, head
	var qNode *ListNode
	for pNode != end {
		qNode = pNode.Next
		pNode.Next = ret
		ret = pNode
		pNode = qNode
	}
	return ret
}
