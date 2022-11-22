package medium

import "math"

func NumSquares(n int) int {
	minForEveryNumBeforeN := []int{0, 1}

	for i := 2; i <= n; i++ {
		minForEveryNumBeforeN = append(minForEveryNumBeforeN, int(math.Pow(10, 4)))
		for j := 1; i >= j*j; j++ {
			minForEveryNumBeforeN[i] = int(math.Min(float64(minForEveryNumBeforeN[i]), float64(minForEveryNumBeforeN[i-j*j])))
		}
		minForEveryNumBeforeN[i]++
	}

	return minForEveryNumBeforeN[n]
}
