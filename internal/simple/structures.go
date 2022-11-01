package simple

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(elements ...int) *ListNode {
	slice := make([]ListNode, len(elements), len(elements))

	for i, el := range elements {
		slice[i] = ListNode{
			Val: el,
		}
	}

	for index := range slice {
		if len(slice) > index+1 {
			slice[index].Next = &slice[index+1]
		}
	}

	return &slice[0]
}
