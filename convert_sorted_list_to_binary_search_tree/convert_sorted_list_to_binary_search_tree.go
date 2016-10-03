package convert_sorted_list_to_binary_search_tree

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	n := 0
	tmp := head
	for tmp != nil {
		n++
		tmp = tmp.Next
	}
	return helper(&head, 0, n-1)
}

func helper(head **ListNode, i, j int) *TreeNode {
	if i > j || *head == nil {
		return nil
	}
	mid := i + (j-i)/2
	left := helper(head, i, mid-1)
	root := &TreeNode{Val: (*head).Val}
	root.Left = left
	*head = (*head).Next
	root.Right = helper(head, mid+1, j)
	return root
}
