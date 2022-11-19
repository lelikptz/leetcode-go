package hard

import (
	"fmt"
	"sort"
)

func OuterTrees(trees [][]int) [][]int {
	sort.Slice(trees, func(p, q int) bool {
		if trees[q][0]-trees[p][0] == 0 {
			return trees[q][1]-trees[p][1] > 0
		} else {
			return trees[q][0]-trees[p][0] > 0
		}
	})

	var hull [][]int
	for i := 0; i < len(trees); i++ {
		for len(hull) >= 2 && orientation(hull[len(hull)-2], hull[len(hull)-1], trees[i]) > 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, trees[i])
	}
	hull = hull[:len(hull)-1]
	fmt.Println(hull)
	for i := len(trees) - 1; i >= 0; i-- {
		for len(hull) >= 2 && orientation(hull[len(hull)-2], hull[len(hull)-1], trees[i]) > 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, trees[i])
	}

	return unique(hull)
}

func unique(hull [][]int) [][]int {

	un := make(map[[2]int][2]int)
	for i := 0; i < len(hull); i++ {
		un[[2]int{hull[i][0], hull[i][1]}] = [2]int{hull[i][0], hull[i][1]}
	}

	v := make([][]int, 0, len(un))
	for k := range un {
		v = append(v, []int{k[0], k[1]})
	}

	return v
}

func orientation(p []int, q []int, r []int) int {
	return (q[1]-p[1])*(r[0]-q[0]) - (q[0]-p[0])*(r[1]-q[1])
}
