package binary_tree_level_order_traversal_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	tmpRets := [][]int{}
	if root == nil {
		return tmpRets
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		tmps := []*TreeNode{}
		ret := []int{}
		for i := range stack {
			ret = append(ret, stack[i].Val)
			if stack[i].Left != nil {
				tmps = append(tmps, stack[i].Left)
			}
			if stack[i].Right != nil {
				tmps = append(tmps, stack[i].Right)
			}
		}
		tmpRets = append(tmpRets, ret)
		stack = tmps
	}
	rets := make([][]int, len(tmpRets))
	for i := range tmpRets {
		rets[len(tmpRets)-1-i] = tmpRets[i]
	}
	return rets
}
