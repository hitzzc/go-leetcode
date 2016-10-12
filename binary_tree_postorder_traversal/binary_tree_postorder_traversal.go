package binary_tree_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	rets := []int{}
	if root == nil {
		return rets
	}
	stack := []*TreeNode{root}
	head := root
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		if node.Left == head || node.Right == head || node.Left == nil && node.Right == nil {
			rets = append(rets, node.Val)
			stack = stack[:len(stack)-1]
			head = node
		} else {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}
	return rets
}
