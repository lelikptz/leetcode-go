package easy

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"math"
)

func HasCycle(head *structure.ListNode) bool {
	max := int(math.Pow(10.0, 5.0))
	cur := head
	for {
		if cur == nil {
			return false
		}
		if cur.Val == max {
			return true
		}
		cur.Val = max
		cur = cur.Next
	}
}
