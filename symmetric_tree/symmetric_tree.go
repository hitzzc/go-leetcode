package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right != nil || root.Left != nil && root.Right == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	leftStack := []*TreeNode{root.Left}
	rightStack := []*TreeNode{root.Right}
	for len(leftStack) != 0 {
		leftStackTmp := []*TreeNode{}
		rightStackTmp := []*TreeNode{}
		for i := range leftStack {
			if leftStack[i].Val != rightStack[i].Val {
				return false
			}

			if leftStack[i].Left == nil && rightStack[i].Right != nil || leftStack[i].Left != nil && rightStack[i].Right == nil {
				return false
			}
			if leftStack[i].Left != nil && rightStack[i].Right != nil {
				leftStackTmp = append(leftStackTmp, leftStack[i].Left)
				rightStackTmp = append(rightStackTmp, rightStack[i].Right)
			}

			if leftStack[i].Right == nil && rightStack[i].Left != nil || leftStack[i].Right != nil && rightStack[i].Left == nil {
				return false
			}
			if leftStack[i].Right != nil && rightStack[i].Left != nil {
				leftStackTmp = append(leftStackTmp, leftStack[i].Right)
				rightStackTmp = append(rightStackTmp, rightStack[i].Left)
			}
		}
		leftStack = leftStackTmp
		rightStack = rightStackTmp
	}
	return true
}
