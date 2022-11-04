package simple

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
)

func InorderTraversal(root *structure.TreeNode) []int {
	var r []int
	queue := make([]structure.TreeNode, 0)

	node := root
	for node != nil || len(queue) > 0 {
		for node != nil {
			queue = append(queue, *node)
			node = node.Left
		}
		node = &queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]
		r = append(r, node.Val)
		node = node.Right
	}

	return r
}

func InorderTraversalRecurse(root *structure.TreeNode) []int {
	var r []int

	return visit(r, root)
}

func visit(r []int, node *structure.TreeNode) []int {
	if nil == node {
		return r
	}

	r = visit(r, node.Left)
	r = append(r, node.Val)
	r = visit(r, node.Right)

	return r
}
