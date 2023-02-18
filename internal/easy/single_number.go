package easy

func SingleNumber(nums []int) int {
	var m = make(map[int]int, len(nums))
	for _, value := range nums {
		m[value]++
	}

	for index, value := range m {
		if value == 1 {
			return index
		}
	}

	return 0
}
