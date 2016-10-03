package minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left, right := -1, -1
	if root.Left != nil {
		left = minDepth(root.Left)
	}
	if root.Right != nil {
		right = minDepth(root.Right)
	}
	depth := 1
	if left == -1 {
		depth += right
	} else if right == -1 {
		depth += left
	} else if left < right {
		depth += left
	} else {
		depth += right
	}
	return depth
}
