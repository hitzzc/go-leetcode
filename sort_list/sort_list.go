package sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	secHead := split(head)
	return merge(sortList(head), sortList(secHead))
}

func split(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	var tail *ListNode
	for fast != nil && fast.Next != nil {
		tail = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	tail.Next = nil
	return slow
}

func merge(left, right *ListNode) (head *ListNode) {
	if left == nil {
		return right
	} else if right == nil {
		return left
	}
	var p, cur *ListNode
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur = left
			left = left.Next
		} else {
			cur = right
			right = right.Next
		}
		if head == nil {
			head = cur
		} else {
			p.Next = cur
		}
		p = cur
	}
	if left != nil {
		p.Next = left
	} else {
		p.Next = right
	}
	return head
}
