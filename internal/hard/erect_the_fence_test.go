package hard

import (
	"testing"
)

func TestOuterTrees(t *testing.T) {
	type args struct {
		trees [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "First",
			args: args{
				trees: [][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}},
			},
			want: [][]int{{1, 1}, {2, 0}, {3, 3}, {2, 4}, {4, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OuterTrees(tt.args.trees)
			if len(got) != len(tt.want) {
				t.Errorf("OuterTrees() = %v, want %v", got, tt.want)
			}

			for _, val := range got {
				if !sliceHasElement(tt.want, val) {
					t.Errorf("OuterTrees() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func sliceHasElement(sl [][]int, find []int) bool {
	for _, val := range sl {
		if val[0] == find[0] && val[1] == find[1] {
			return true
		}
	}

	return false
}
