package medium

import (
	"math"
)

func ComputeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	return (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1) - common(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
}

func common(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	width := math.Min(float64(ax2), float64(bx2)) - math.Max(float64(ax1), float64(bx1))
	height := math.Min(float64(ay2), float64(by2)) - math.Max(float64(ay1), float64(by1))
	if width < 0 || height < 0 {
		return 0
	}

	return int(width * height)
}
