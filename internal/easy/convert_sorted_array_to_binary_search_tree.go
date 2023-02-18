package easy

import "github.com/lelikptz/leetcode-go/internal/structure"

func SortedArrayToBST(nums []int) *structure.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	middle := nums[len(nums)/2]

	return &structure.TreeNode{
		Val:   middle,
		Left:  SortedArrayToBST(nums[:len(nums)/2]),
		Right: SortedArrayToBST(nums[len(nums)/2+1:]),
	}
}
