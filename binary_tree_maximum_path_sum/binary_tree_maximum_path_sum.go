package binary_tree_maximum_path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type solution struct {
	max   int
	first bool
}

func maxPathSum(root *TreeNode) int {
	s := &solution{first: true}
	s.helper(root)
	return s.max
}

func (this *solution) helper(root *TreeNode) (max int) {
	if root == nil {
		return 0
	}
	tmp := root.Val
	left := this.helper(root.Left)
	if left > 0 {
		tmp += left
	}
	right := this.helper(root.Right)
	if right > 0 {
		tmp += right
	}
	if this.first || tmp > this.max {
		this.first = false
		this.max = tmp
	}
	if left > 0 && left > right {
		max = root.Val + left
	} else if right > 0 && right > left {
		max = root.Val + right
	} else {
		max = root.Val
	}
	return
}
