package easy

func ClimbStairs(n int) int {
	fib := make(map[int]int)
	fib[1] = 1
	fib[2] = 2
	for i := 3; i <= 45; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}
