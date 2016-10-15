package binary_tree_right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rets := []int{}
	queue := []*TreeNode{root}
	last := root
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
		if top == last {
			rets = append(rets, top.Val)
			if len(queue) != 0 {
				last = queue[len(queue)-1]
			}
		}
	}
	return rets
}
