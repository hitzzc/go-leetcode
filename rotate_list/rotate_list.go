package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	cnt := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		cnt++
	}
	k %= cnt
	if k == 0 {
		return head
	}
	pre, post := head, head
	for i := 0; i < k && post != nil; i++ {
		post = post.Next
	}
	if post == nil {
		return head
	}
	for post.Next != nil {
		pre, post = pre.Next, post.Next
	}
	newHead := pre.Next
	post.Next = head
	pre.Next = nil
	return newHead
}
