package simple

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
)

type Item struct {
	Node *structure.TreeNode
	Sum  int
}

func HasPathSum(root *structure.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	queue := []Item{{Node: root, Sum: root.Val}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Node.Right == nil && current.Node.Left == nil && current.Sum == targetSum {
			return true
		}
		if current.Node.Right != nil {
			queue = append(queue, Item{
				Node: current.Node.Right,
				Sum:  current.Node.Right.Val + current.Sum,
			})
		}
		if current.Node.Left != nil {
			queue = append(queue, Item{
				Node: current.Node.Left,
				Sum:  current.Node.Left.Val + current.Sum,
			})
		}
	}

	return false
}
