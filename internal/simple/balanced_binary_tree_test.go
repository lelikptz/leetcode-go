package simple

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	type args struct {
		root *structure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "First",
			args: args{
				root: nil,
			},
			want: true,
		},
		{
			name: "Second",
			args: args{
				root: &structure.TreeNode{
					Val: 1,
					Left: &structure.TreeNode{
						Val: 2,
						Left: &structure.TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &structure.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: true,
		},
		{
			name: "Third",
			args: args{
				root: &structure.TreeNode{
					Val: 1,
					Left: &structure.TreeNode{
						Val: 2,
						Left: &structure.TreeNode{
							Val: 3,
							Left: &structure.TreeNode{
								Val:   4,
								Left:  nil,
								Right: nil,
							},
							Right: nil,
						},
						Right: nil,
					},
					Right: &structure.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBalanced(tt.args.root); got != tt.want {
				t.Errorf("IsBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
