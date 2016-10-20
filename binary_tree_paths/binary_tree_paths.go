package binary_tree_paths

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type solution struct {
	path []*TreeNode
	ret  []string
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	s := &solution{path: []*TreeNode{root}}
	s.helper()
	return s.ret
}

func (this *solution) helper() {
	if len(this.path) == 0 {
		return
	}
	root := this.path[len(this.path)-1]
	if root.Left == nil && root.Right == nil {
		s := strconv.Itoa(this.path[0].Val)
		for i := 1; i < len(this.path); i++ {
			s += "->" + strconv.Itoa(this.path[i].Val)
		}
		this.ret = append(this.ret, s)
		return
	}
	if root.Left != nil {
		this.path = append(this.path, root.Left)
		this.helper()
		this.path = this.path[:len(this.path)-1]
	}
	if root.Right != nil {
		this.path = append(this.path, root.Right)
		this.helper()
		this.path = this.path[:len(this.path)-1]
	}
	return
}
