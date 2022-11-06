package simple

import "github.com/lelikptz/leetcode-go/internal/structure"

func IsSameTree(p *structure.TreeNode, q *structure.TreeNode) bool {
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}

	return p == nil && q == nil || (p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right))
}
