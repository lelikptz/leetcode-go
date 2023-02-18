package easy

func GetRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	prev := GetRow(rowIndex - 1)
	next := []int{1}

	for i := 1; i < len(prev); i++ {
		next = append(next, prev[i]+prev[i-1])
	}
	next = append(next, 1)

	return next
}
