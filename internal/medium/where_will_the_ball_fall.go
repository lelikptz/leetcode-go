package medium

func FindBall(grid [][]int) []int {
	res := make([]int, 0, len(grid[0]))
	for column := 0; column < len(grid[0]); column++ {
		start := column
		for row := range grid {
			if canLeft(grid, row, start) {
				start--
			} else if canRight(grid, row, start) {
				start++
			} else {
				start = -1
				break
			}
		}
		res = append(res, start)
	}

	return res
}

func canRight(grid [][]int, row int, column int) bool {
	if len(grid[row]) <= column+1 {
		return false
	}
	if grid[row][column] == 1 && grid[row][column+1] == 1 {
		return true
	}

	return false
}

func canLeft(grid [][]int, row int, column int) bool {
	if column < 1 {
		return false
	}
	if grid[row][column] == -1 && grid[row][column-1] == -1 {
		return true
	}

	return false
}
