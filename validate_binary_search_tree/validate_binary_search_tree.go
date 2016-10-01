package validate_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	p := root
	stack := []*TreeNode{}
	sortedValues := []int{}
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(sortedValues) != 0 && p.Val <= sortedValues[len(sortedValues)-1] {
				return false
			}
			sortedValues = append(sortedValues, p.Val)
			p = p.Right
		}
	}
	return true
}
