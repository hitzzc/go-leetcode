package unique_binary_search_trees_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTreesI(1, n)
}

func generateTreesI(left, right int) (ret []*TreeNode) {
	if left > right {
		return []*TreeNode{nil}
	}
	if left == right {
		return []*TreeNode{&TreeNode{Val: left}}
	}
	for num := left; num <= right; num++ {
		leftTrees := generateTreesI(left, num-1)
		rightTrees := generateTreesI(num+1, right)
		for i := range leftTrees {
			for j := range rightTrees {
				root := &TreeNode{Val: num}
				root.Left = leftTrees[i]
				root.Right = rightTrees[j]
				ret = append(ret, root)
			}
		}
	}
	return
}
