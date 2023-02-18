package easy

import "github.com/lelikptz/leetcode-go/internal/structure"

func MergeTwoLists(a *structure.ListNode, b *structure.ListNode) *structure.ListNode {
	head := &structure.ListNode{
		Val:  0,
		Next: nil,
	}

	cur := head
	for a != nil && b != nil {
		if a.Val > b.Val {
			cur.Next = b
			b = b.Next
		} else {
			cur.Next = a
			a = a.Next
		}
		cur = cur.Next
	}
	if a != nil {
		cur.Next = a
	} else {
		cur.Next = b
	}

	return head.Next
}
