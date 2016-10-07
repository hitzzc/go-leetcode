package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var pre *TreeNode
	helper(root, &pre)
	return
}

func helper(root *TreeNode, pre **TreeNode) {
	if root == nil {
		return
	}
	if *pre != nil {
		(*pre).Left = nil
		(*pre).Right = root
	}
	*pre = root
	left := root.Left
	right := root.Right
	if left != nil {
		helper(left, pre)
	}
	if right != nil {
		helper(right, pre)
	}
	return
}
