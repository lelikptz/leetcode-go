package main

import (
	"bitbucket.org/lelik-orlov/leetcode/internal/problems"
	"fmt"
)

func main() {

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	problems.RemoveDuplicates(nums)
	fmt.Println(nums)
}
