package simple

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	type args struct {
		p *structure.TreeNode
		q *structure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "First",
			args: args{
				p: &structure.TreeNode{
					Val:   1,
					Left:  &structure.TreeNode{Val: 2, Left: &structure.TreeNode{Val: 4}},
					Right: &structure.TreeNode{Val: 3},
				},
				q: &structure.TreeNode{
					Val:   1,
					Left:  &structure.TreeNode{Val: 2, Left: &structure.TreeNode{Val: 4}},
					Right: &structure.TreeNode{Val: 3},
				},
			},
			want: true,
		},
		{
			name: "Second",
			args: args{
				p: &structure.TreeNode{
					Val:   1,
					Left:  &structure.TreeNode{Val: 2, Left: &structure.TreeNode{Val: 4}},
					Right: &structure.TreeNode{Val: 3},
				},
				q: &structure.TreeNode{
					Val:   1,
					Left:  &structure.TreeNode{Val: 12, Left: &structure.TreeNode{Val: 4}},
					Right: &structure.TreeNode{Val: 3},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("IsSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
