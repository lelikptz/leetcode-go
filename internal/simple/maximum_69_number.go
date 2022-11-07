package simple

import (
	"strconv"
)

func Maximum69Number(num int) int {
	max := num
	s := []rune(strconv.Itoa(num))
	nine := []rune("96")
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "6" {
			s[i] = nine[0]
			n, _ := strconv.Atoi(string(s))
			if n > max {
				max = n
			}
			s[i] = nine[1]
		}
	}

	return max
}
