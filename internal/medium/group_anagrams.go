package medium

import (
	"sort"
	"strings"
)

func GroupAnagrams(sliceStrings []string) [][]string {
	m := map[string][]string{}
	for _, s := range sliceStrings {
		sl := strings.Split(s, "")
		sort.Strings(sl)
		key := strings.Join(sl, "")
		val, ok := m[key]
		if ok {
			m[key] = append(val, s)
		} else {
			m[key] = []string{s}
		}
	}

	var res [][]string
	for _, s := range m {
		res = append(res, s)
	}

	return res
}
