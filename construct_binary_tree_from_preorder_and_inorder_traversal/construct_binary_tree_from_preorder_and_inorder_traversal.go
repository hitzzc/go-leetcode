package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := map[int]int{}
	for i := range inorder {
		m[inorder[i]] = i
	}
	return helper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, m)
}

func helper(preorder []int, preL, preR int, inorder []int, inL, inR int, m map[int]int) *TreeNode {
	if preL > preR || inL > inR {
		return nil
	}
	root := &TreeNode{Val: preorder[preL]}
	index := m[root.Val]
	root.Left = helper(preorder, preL+1, preL+index-inL, inorder, inL, index-1, m)
	root.Right = helper(preorder, preL+index-inL+1, preR, inorder, index+1, inR, m)
	return root
}
