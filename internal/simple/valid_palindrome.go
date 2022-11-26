package simple

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	regex, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = strings.ToLower(regex.ReplaceAllString(s, ""))
	stringLen := len(s)
	iMax := stringLen / 2

	for i := 0; i < iMax; i++ {
		if s[i] != s[stringLen-i-1] {
			return false
		}
	}

	return true
}
