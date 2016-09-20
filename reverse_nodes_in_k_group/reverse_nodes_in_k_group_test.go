package reverse_nodes_in_k_group

import (
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	reverseKGroup(head, 2)
}
