package medium

import "sort"

func LongestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	cache := make(map[string]int)

	sum := 0
	for i := 0; i < len(words); i++ {
		maxLevel := 0
		createTree(words[i], words[i+1:], 0, &maxLevel, cache)
		if maxLevel > sum {
			sum = maxLevel
		}
	}

	return sum
}

func createTree(root string, words []string, level int, maxLevel *int, cache map[string]int) {
	level++
	if level > *maxLevel {
		*maxLevel = level
	}
	i := 0
	for i < len(words) {
		if predecessor(root, words[i]) {
			if val, ok := cache[words[i]]; ok {
				*maxLevel = val
			} else {
				createTree(words[i], words[1:], level, maxLevel, cache)
				cache[words[i]] = *maxLevel
			}
		}
		i++
	}
	return
}

func predecessor(a, b string) bool {
	if len(b)-len(a) != 1 {
		return false
	}

	var counter int
	for i, j := 0, 0; i < len(a) && j < len(b); {
		if a[i] == b[j] {
			counter++
			i++
		}
		j++
	}

	return counter == len(a)
}
