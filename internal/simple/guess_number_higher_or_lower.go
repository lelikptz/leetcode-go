package simple

import (
	"math"
)

func GuessNumber(n int) int {
	sl := calculate(n, 1, int(math.Pow(2, 31)))

	for val := guess(sl[0]); val != 0; val = guess(sl[0]) {
		sl = calculate(sl[0], sl[1], sl[2])
	}

	return sl[0]
}

func calculate(n int, min int, max int) []int {
	if guess(n) > 0 {
		min = n
	}

	if guess(n) < 0 {
		max = n
	}
	return []int{(min + max) / 2, min, max}
}

var GuessValue int

func guess(num int) int {
	if num == GuessValue {
		return 0
	}

	if num > GuessValue {
		return -1
	} else {
		return 1
	}
}
