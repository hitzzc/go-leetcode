package palindrome_lined_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	head2 := reverse(slow)
	for ; head != nil && head2 != nil; head, head2 = head.Next, head2.Next {
		if head.Val != head2.Val {
			return false
		}
	}
	return true
}

func reverse(head *ListNode) (newHead *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		if tmp == nil {
			break
		}
		cur = tmp
	}
	return cur
}
