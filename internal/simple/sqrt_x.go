package simple

func MySqrt(x int) int {
	i := 0
	for i*i <= x {
		i++
	}
	if x < i*i {
		i--
	}
	return i
}
