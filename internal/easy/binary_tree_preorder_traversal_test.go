package easy

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	type args struct {
		root *structure.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "First",
			args: args{
				root: &structure.TreeNode{Val: 3, Left: &structure.TreeNode{Val: 1, Left: nil, Right: &structure.TreeNode{Val: 2}}, Right: nil},
			},
			want: []int{3, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
