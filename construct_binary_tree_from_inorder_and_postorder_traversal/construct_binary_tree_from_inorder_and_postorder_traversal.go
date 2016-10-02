package construct_binary_tree_from_inorder_and_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	m := map[int]int{}
	for i := range inorder {
		m[inorder[i]] = i
	}
	return helper(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1, m)
}

func helper(inorder []int, inL, inR int, postorder []int, postL, postR int, m map[int]int) *TreeNode {
	if inL > inR || postL > postR {
		return nil
	}
	root := &TreeNode{Val: postorder[postR]}
	index := m[root.Val]
	root.Left = helper(inorder, inL, index-1, postorder, postL, postL+index-inL-1, m)
	root.Right = helper(inorder, index+1, inR, postorder, postL+index-inL, postR-1, m)
	return root
}
