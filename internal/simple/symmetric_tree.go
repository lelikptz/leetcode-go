package simple

import "github.com/lelikptz/leetcode-go/internal/structure"

func IsSymmetric(root *structure.TreeNode) bool {
	return checkIsSymmetricTree(root.Left, root.Right)
}

func checkIsSymmetricTree(left *structure.TreeNode, right *structure.TreeNode) bool {
	if left == nil && right != nil || left != nil && right == nil {
		return false
	}

	return (left == nil && right == nil) || (left.Val == right.Val && checkIsSymmetricTree(left.Left, right.Right) && checkIsSymmetricTree(left.Right, right.Left))
}
