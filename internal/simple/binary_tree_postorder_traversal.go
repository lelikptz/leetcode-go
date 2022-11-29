package simple

import "github.com/lelikptz/leetcode-go/internal/structure"

func PostorderTraversal(root *structure.TreeNode) []int {
	var r []int

	return visitPostorder(r, root)
}

func visitPostorder(r []int, node *structure.TreeNode) []int {
	if nil == node {
		return r
	}

	r = visitPostorder(r, node.Left)
	r = visitPostorder(r, node.Right)
	r = append(r, node.Val)

	return r
}
