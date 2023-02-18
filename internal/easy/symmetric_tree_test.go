package easy

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
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
				root: &structure.TreeNode{
					Val:   1,
					Left:  &structure.TreeNode{Val: 2, Left: &structure.TreeNode{Val: 3}, Right: &structure.TreeNode{Val: 4}},
					Right: &structure.TreeNode{Val: 2, Left: &structure.TreeNode{Val: 4}, Right: &structure.TreeNode{Val: 3}},
				},
			},
			want: true,
		},
		{
			name: "Second",
			args: args{
				root: &structure.TreeNode{
					Val:   1,
					Left:  &structure.TreeNode{Val: 2, Left: &structure.TreeNode{Val: 3}, Right: &structure.TreeNode{Val: 3}},
					Right: &structure.TreeNode{Val: 2, Left: &structure.TreeNode{Val: 4}, Right: &structure.TreeNode{Val: 4}},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSymmetric(tt.args.root); got != tt.want {
				t.Errorf("IsSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
