package simple

func Generate(numRows int) [][]int {
	var triangle [][]int
	var row []int

	for i := 1; i <= numRows; i++ {
		row = []int{}
		for j := 1; j <= i; j++ {
			if i >= 2 && j >= 2 && len(triangle[i-2]) > j-1 {
				row = append(row, triangle[i-2][j-2]+triangle[i-2][j-1])
			} else {
				row = append(row, 1)
			}
		}

		triangle = append(triangle, row)
	}

	return triangle
}
