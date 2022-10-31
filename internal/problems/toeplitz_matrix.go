package problems

func IsToeplitzMatrix(matrix [][]int) bool {
	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[0]); column++ {
			if len(matrix) >= row+1 && len(matrix[0]) >= column+1 {
				if row+1 < len(matrix) && column+1 < len(matrix[0]) {
					if matrix[row][column] != matrix[row+1][column+1] {
						return false
					}
				}
			}
		}
	}

	return true
}
