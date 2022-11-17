package simple

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"math"
)

func IsBalanced(root *structure.TreeNode) bool {
	if root == nil {
		return true
	}

	return math.Abs(float64(getDepth(root.Left)-getDepth(root.Right))) <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}

func getDepth(node *structure.TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + int(math.Max(float64(getDepth(node.Left)), float64(getDepth(node.Right))))
}
