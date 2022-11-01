package simple

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	head := &ListNode{
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
