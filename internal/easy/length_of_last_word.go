package easy

import (
	"strings"
)

func LengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")

	counter := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			return counter
		}
		counter++
	}

	return counter
}
