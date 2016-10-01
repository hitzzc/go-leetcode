package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}
	rets := [][]int{}
	for len(stack) != 0 {
		tmpStack := []*TreeNode{}
		ret := []int{}
		for i := range stack {
			ret = append(ret, stack[i].Val)
			if stack[i].Left != nil {
				tmpStack = append(tmpStack, stack[i].Left)
			}
			if stack[i].Right != nil {
				tmpStack = append(tmpStack, stack[i].Right)
			}
		}
		rets = append(rets, ret)
		stack = tmpStack
	}
	return rets
}
