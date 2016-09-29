package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	p := root
	ret := []int{}
	stack := []*TreeNode{}
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = tmp.Right
			ret = append(ret, tmp.Val)
		}
	}
	return ret
}
