package medium

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	m := map[[100]string][]string{}
	for _, s := range strs {
		sl := strings.Split(s, "")
		sort.Strings(sl)
		key := [100]string{strings.Join(sl, "")}
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
