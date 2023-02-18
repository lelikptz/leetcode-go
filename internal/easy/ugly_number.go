package easy

func IsUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for _, val := range []int{2, 3, 5} {
		for n%val == 0 {
			n = n / val
		}
	}

	return n == 1
}
