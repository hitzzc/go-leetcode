package kth_smallest_element_in_a_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	if left == k-1 {
		return root.Val
	}
	if left > k-1 {
		return kthSmallest(root.Left, k)
	}
	return kthSmallest(root.Right, k-1-left)
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + helper(root.Left) + helper(root.Right)
}
