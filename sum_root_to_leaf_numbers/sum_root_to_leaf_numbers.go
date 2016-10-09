package sum_root_to_leaf_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type solution struct {
	sum int
}

func sumNumbers(root *TreeNode) int {
	s := &solution{}
	s.helper(root, 0)
	return s.sum
}

func (this *solution) helper(root *TreeNode, current int) {
	if root == nil {
		return
	}
	v := current*10 + root.Val
	if root.Left == nil && root.Right == nil {
		this.sum += v
	}
	this.helper(root.Left, v)
	this.helper(root.Right, v)
	return
}
