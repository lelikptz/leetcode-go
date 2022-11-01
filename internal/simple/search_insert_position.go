package simple

func searchInsert(nums []int, target int) int {
	for index, value := range nums {
		if value >= target {
			return index
		}
	}

	return len(nums)
}
