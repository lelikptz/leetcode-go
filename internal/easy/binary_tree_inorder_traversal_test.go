package easy

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"testing"
)

type args struct {
	root *structure.TreeNode
}
type inorderTraversal struct {
	name string
	args args
	want []int
}

func TestInorderTraversal(t *testing.T) {
	tests := getTestDataForBinaryTreeInorderTraversal()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InorderTraversal(tt.args.root)
			if len(got) != len(tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}

			for i := 1; i < len(tt.want); i++ {
				if tt.want[i] != got[i] {
					t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestInorderTraversalRecurse(t *testing.T) {
	tests := getTestDataForBinaryTreeInorderTraversal()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InorderTraversalRecurse(tt.args.root)
			if len(got) != len(tt.want) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}

			for i := 1; i < len(tt.want); i++ {
				if tt.want[i] != got[i] {
					t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func getTestDataForBinaryTreeInorderTraversal() []inorderTraversal {
	return []inorderTraversal{
		{
			name: "First",
			args: args{
				root: &structure.TreeNode{
					Val:  1,
					Left: nil,
					Right: &structure.TreeNode{
						Val: 2,
						Left: &structure.TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "Second",
			args: args{
				root: nil,
			},
			want: nil,
		},
		{
			name: "Third",
			args: args{
				root: &structure.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: []int{1},
		},
		{
			name: "Fourth",
			args: args{
				root: &structure.TreeNode{
					Val: 2,
					Left: &structure.TreeNode{
						Val: 3,
						Left: &structure.TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: nil,
				},
			},
			want: []int{1, 3, 2},
		},
	}
}
