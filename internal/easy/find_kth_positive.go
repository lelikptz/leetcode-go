package easy

func FindKthPositive(arr []int, k int) int {
	count := 0
	j := 0
	for i := 1; i <= len(arr)+k; i++ {
		if j >= len(arr) || arr[j] != i {
			count++
		} else {
			j++
		}
		if count == k {
			return i
		}
	}

	return 0
}
