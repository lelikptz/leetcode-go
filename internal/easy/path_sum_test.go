package easy

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	type args struct {
		root      *structure.TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "First",
			args: args{
				root: &structure.TreeNode{
					Val: 5,
					Left: &structure.TreeNode{
						Val: 4,
						Left: &structure.TreeNode{
							Val:   11,
							Left:  &structure.TreeNode{Val: 7},
							Right: &structure.TreeNode{Val: 2},
						},
						Right: nil,
					},
					Right: &structure.TreeNode{
						Val: 8,
						Left: &structure.TreeNode{
							Val:   13,
							Left:  nil,
							Right: nil,
						},
						Right: &structure.TreeNode{
							Val:  4,
							Left: nil,
							Right: &structure.TreeNode{
								Val:   1,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "Second",
			args: args{
				root: &structure.TreeNode{
					Val:   1,
					Left:  &structure.TreeNode{Val: 3},
					Right: &structure.TreeNode{Val: 2},
				},
				targetSum: 5,
			},
			want: false,
		},
		{
			name: "Third",
			args: args{
				root: &structure.TreeNode{
					Val:  1,
					Left: &structure.TreeNode{Val: 2},
				},
				targetSum: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("HasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
