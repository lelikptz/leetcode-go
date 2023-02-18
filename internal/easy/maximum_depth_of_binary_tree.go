package easy

import "github.com/lelikptz/leetcode-go/internal/structure"

func MaxDepth(root *structure.TreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	if left > right {
		return left + 1
	}

	return right + 1
}
