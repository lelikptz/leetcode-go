package medium

import "github.com/lelikptz/leetcode-go/internal/structure"

func CountNodes(root *structure.TreeNode) int {
	var counter = 0
	return Count(root, &counter)
}

func Count(root *structure.TreeNode, counter *int) int {
	if root != nil {
		*counter++
		Count(root.Left, counter)
		Count(root.Right, counter)
	}

	return *counter
}
