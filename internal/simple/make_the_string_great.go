package simple

import "strings"

func MakeGood(s string) string {
	sl := []rune(s)
	for i := 0; i < len(sl)-1; i++ {
		if isBad(sl[i], sl[i+1]) {
			return MakeGood(string(sl[:i]) + string(sl[i+2:]))
		}
	}

	return s
}

func isBad(p rune, q rune) bool {
	return (string(p) == strings.ToLower(string(p)) && strings.ToUpper(string(p)) == string(q)) || (string(p) == strings.ToUpper(string(p)) && strings.ToLower(string(p)) == string(q))
}
