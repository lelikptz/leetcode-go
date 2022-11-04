package simple

import "github.com/lelikptz/leetcode-go/internal/structure"

func DeleteDuplicates(head *structure.ListNode) *structure.ListNode {
	cur := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return cur
}
