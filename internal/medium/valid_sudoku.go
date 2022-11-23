package medium

func IsValidSudoku(board [][]byte) bool {
	var column []byte
	cubes := make([][]byte, 9, 9)

	for i := 0; i < 9; i++ {
		if !isValid(board[i]) {
			return false
		}
		column = []byte{}
		for j := 0; j < 9; j++ {
			column = append(column, board[j][i])
			cubes[i/3+j/3*3] = append(cubes[i/3+j/3*3], board[i][j])
		}
		if !isValid(column) {
			return false
		}
	}

	for _, val := range cubes {
		if !isValid(val) {
			return false
		}
	}

	return true
}

func isValid(nums []byte) bool {
	unique := make(map[byte]int, 9)
	for _, b := range []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'} {
		unique[b] = 0
	}

	for _, element := range nums {
		if element != '.' {
			unique[element]++

			val, _ := unique[element]
			if val > 1 {
				return false
			}

		}
	}

	return true
}
