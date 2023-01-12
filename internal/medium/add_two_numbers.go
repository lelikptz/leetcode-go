package medium

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
)

func AddTwoNumbers(l1 *structure.ListNode, l2 *structure.ListNode) *structure.ListNode {
	result := &structure.ListNode{}
	carry := 0
	cur := result
	for l1 != nil || l2 != nil {
		var curVal int
		if l1 == nil {
			curVal = l2.Val + carry
		} else if l2 == nil {
			curVal = l1.Val + carry
		} else {
			curVal = l1.Val + l2.Val + carry
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if curVal >= 10 {
			carry = curVal / 10
			curVal %= 10
		} else {
			carry = 0
		}

		cur.Val = curVal
		if l1 != nil || l2 != nil {
			cur.Next = &structure.ListNode{}
			cur = cur.Next
		}
	}

	if carry > 0 {
		cur.Next = &structure.ListNode{Val: carry}
	}

	return result
}
