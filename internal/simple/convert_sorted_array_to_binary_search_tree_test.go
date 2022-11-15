package simple

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *structure.TreeNode
	}{
		{
			name: "First",
			args: args{
				nums: []int{1, 3},
			},
			want: &structure.TreeNode{
				Val: 3,
				Left: &structure.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		}, {
			name: "Second",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: &structure.TreeNode{
				Val: 0,
				Left: &structure.TreeNode{
					Val: -3,
					Left: &structure.TreeNode{
						Val:   -10,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: &structure.TreeNode{
					Val: 9,
					Left: &structure.TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
