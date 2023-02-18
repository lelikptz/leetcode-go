package easy

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	type args struct {
		root *structure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First",
			args: args{
				root: nil,
			},
			want: 0,
		},
		{
			name: "Second",
			args: args{
				root: &structure.TreeNode{Val: 1},
			},
			want: 1,
		},
		{
			name: "Third",
			args: args{
				root: &structure.TreeNode{
					Val: 1,
					Right: &structure.TreeNode{Val: 2,
						Left: &structure.TreeNode{Val: 3,
							Left: &structure.TreeNode{Val: 4},
						},
					},
				},
			},
			want: 4,
		},
		{
			name: "Fourth",
			args: args{
				root: &structure.TreeNode{
					Val: 1,
					Left: &structure.TreeNode{Val: 2,
						Left:  &structure.TreeNode{Val: 4},
						Right: &structure.TreeNode{Val: 5},
					},
					Right: &structure.TreeNode{Val: 3,
						Left:  &structure.TreeNode{Val: 6},
						Right: &structure.TreeNode{Val: 7},
					},
				},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDepth(tt.args.root); got != tt.want {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
