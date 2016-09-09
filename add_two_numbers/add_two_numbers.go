package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	tmp1, tmp2 := l1, l2
	head := new(ListNode)
	var walk *ListNode
	first := true
	for tmp1 != nil && tmp2 != nil {
		walk, carry = step(walk, head, first, tmp1.Val+tmp2.Val+carry)
		first = false
		tmp1, tmp2 = tmp1.Next, tmp2.Next
	}
	var remain *ListNode
	if tmp1 != nil {
		remain = tmp1
	} else {
		remain = tmp2
	}
	for remain != nil {
		walk, carry = step(walk, head, first, remain.Val+carry)
		first = false
		remain = remain.Next
	}
	if carry != 0 {
		walk.Next = &ListNode{Val: carry}
	}
	return head
}

func step(walk, head *ListNode, first bool, val int) (newWalk *ListNode, carry int) {
	if first {
		newWalk = head
	} else {
		walk.Next = new(ListNode)
		newWalk = walk.Next
	}
	if val > 9 {
		newWalk.Val = val - 10
		carry = 1
	} else {
		newWalk.Val = val
		carry = 0
	}
	return
}
