package medium

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"testing"
)

func TestCountNodes(t *testing.T) {
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
				root: &structure.TreeNode{Val: 1, Left: &structure.TreeNode{Val: 2}, Right: &structure.TreeNode{Val: 3}},
			},
			want: 3,
		},
		{
			name: "Third",
			args: args{
				root: &structure.TreeNode{Val: 1, Left: &structure.TreeNode{Val: 2, Left: &structure.TreeNode{Val: 3}, Right: &structure.TreeNode{Val: 4}}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNodes(tt.args.root); got != tt.want {
				t.Errorf("CountNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
