package easy

import "github.com/lelikptz/leetcode-go/internal/structure"

func PreorderTraversal(root *structure.TreeNode) []int {
	var r []int
	queue := make([]structure.TreeNode, 0)

	node := root
	for node != nil || len(queue) > 0 {
		for node != nil {
			r = append(r, node.Val)
			queue = append(queue, *node)
			node = node.Left
		}
		node = &queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]
		node = node.Right
	}

	return r
}
