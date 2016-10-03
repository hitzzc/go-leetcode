package balanced_bianry_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, ret := helper(root)
	return ret
}

func helper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, leftRet := helper(root.Left)
	if !leftRet {
		return -1, false
	}
	rightDepth, rightRet := helper(root.Right)
	if !rightRet {
		return -1, false
	}
	if leftDepth > 1+rightDepth || rightDepth > 1+leftDepth {
		return -1, false
	}
	depth := 1 + leftDepth
	if rightDepth > leftDepth {
		depth = 1 + rightDepth
	}
	return depth, true
}
