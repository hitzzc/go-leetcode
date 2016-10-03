package path_sum_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	return helper(root, sum, []int{})
}

func helper(root *TreeNode, sum int, solution []int) (rets [][]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if root.Val != sum {
			return
		}
		tmp := make([]int, len(solution)+1)
		copy(tmp, solution)
		tmp[len(solution)] = sum
		rets = append(rets, tmp)
	}
	solution = append(solution, root.Val)
	leftRets := helper(root.Left, sum-root.Val, solution)
	rightRets := helper(root.Right, sum-root.Val, solution)
	if len(leftRets) != 0 {
		rets = append(rets, leftRets...)
	}
	if len(rightRets) != 0 {
		rets = append(rets, rightRets...)
	}
	return rets
}
