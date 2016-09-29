package reverse_linked_list_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m == n {
		return head
	}
	fake := &ListNode{Next: head}
	tmp := fake
	for i := 0; i < m-1; i++ {
		tmp = tmp.Next
	}
	current, post := tmp.Next, tmp.Next.Next
	tail := current
	count := n - m
	var ptr *ListNode
	for ; count > 0; count-- {
		ptr = post.Next
		post.Next = current
		current = post
		post = ptr
	}
	tail.Next = post
	tmp.Next = current
	return fake.Next
}
