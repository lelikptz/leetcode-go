package easy

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	type args struct {
		root *structure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *structure.TreeNode
	}{
		{
			name: "First",
			args: args{
				root: nil,
			},
			want: nil,
		},
		{
			name: "Second",
			args: args{
				root: &structure.TreeNode{Val: 1},
			},
			want: &structure.TreeNode{Val: 1},
		},
		{
			name: "Third",
			args: args{
				root: &structure.TreeNode{
					Val:   1,
					Left:  &structure.TreeNode{Val: 2},
					Right: &structure.TreeNode{Val: 3},
				},
			},
			want: &structure.TreeNode{
				Val:   1,
				Left:  &structure.TreeNode{Val: 3},
				Right: &structure.TreeNode{Val: 2},
			},
		},
		{
			name: "Fourth",
			args: args{
				root: &structure.TreeNode{
					Val: 1,
					Left: &structure.TreeNode{
						Val:   2,
						Left:  &structure.TreeNode{Val: 4},
						Right: &structure.TreeNode{Val: 5},
					},
					Right: &structure.TreeNode{
						Val:   3,
						Left:  &structure.TreeNode{Val: 6},
						Right: &structure.TreeNode{Val: 7},
					},
				},
			},
			want: &structure.TreeNode{
				Val: 1,
				Left: &structure.TreeNode{
					Val:   3,
					Left:  &structure.TreeNode{Val: 7},
					Right: &structure.TreeNode{Val: 6},
				},
				Right: &structure.TreeNode{
					Val:   2,
					Left:  &structure.TreeNode{Val: 5},
					Right: &structure.TreeNode{Val: 4},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InvertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
