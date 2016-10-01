package binary_tree_zigzag_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}
	rets := [][]int{}
	zigzag := false
	for len(stack) != 0 {
		tmpStack := []*TreeNode{}
		ret := []int{}
		for i := range stack {
			if zigzag {
				ret = append(ret, stack[len(stack)-i-1].Val)
			} else {
				ret = append(ret, stack[i].Val)
			}
			if stack[i].Left != nil {
				tmpStack = append(tmpStack, stack[i].Left)
			}
			if stack[i].Right != nil {
				tmpStack = append(tmpStack, stack[i].Right)
			}
		}
		zigzag = !zigzag
		rets = append(rets, ret)
		stack = tmpStack
	}
	return rets
}
