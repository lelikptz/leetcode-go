package simple

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"testing"
)

func TestMinDepth(t *testing.T) {
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
				root: &structure.TreeNode{
					Val:  3,
					Left: &structure.TreeNode{Val: 9},
					Right: &structure.TreeNode{
						Val:   20,
						Left:  &structure.TreeNode{Val: 15},
						Right: &structure.TreeNode{Val: 20},
					},
				},
			},
			want: 2,
		},
		{
			name: "Second",
			args: args{
				root: &structure.TreeNode{
					Val: 2,
					Right: &structure.TreeNode{
						Val: 3,
						Right: &structure.TreeNode{
							Val: 4,
							Right: &structure.TreeNode{
								Val: 5,
								Right: &structure.TreeNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDepth(tt.args.root); got != tt.want {
				t.Errorf("MinDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
