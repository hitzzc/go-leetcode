package binary_tree_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	rets := []int{}
	if root == nil {
		return rets
	}
	stack := []*TreeNode{}
	p := root
	for len(stack) != 0 || p != nil {
		for p != nil {
			rets = append(rets, p.Val)
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
	return rets
}
