package easy

import "github.com/lelikptz/leetcode-go/internal/structure"

func InvertTree(root *structure.TreeNode) *structure.TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)

	return root
}
