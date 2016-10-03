package convert_sorted_array_to_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, i, j int) *TreeNode {
	if len(nums) == 0 || i > j || i < 0 || j >= len(nums) {
		return nil
	}
	if i == j {
		return &TreeNode{Val: nums[i]}
	}
	mid := i + (j-i)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, i, mid-1)
	root.Right = helper(nums, mid+1, j)
	return root
}
