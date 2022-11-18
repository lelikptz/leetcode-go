package simple

import "github.com/lelikptz/leetcode-go/internal/structure"

func MinDepth(root *structure.TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	queue := []*structure.TreeNode{root}
	for len(queue) > 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Right == nil && cur.Left == nil {
				return depth
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
		}

		depth++
	}

	return depth
}
