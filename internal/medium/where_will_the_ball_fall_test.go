package medium

import (
	"reflect"
	"testing"
)

func Test_findBall(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "First",
			args: args{
				grid: [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}},
			},
			want: []int{1, -1, -1, -1, -1},
		},
		{
			name: "Second",
			args: args{
				grid: [][]int{{-1}},
			},
			want: []int{-1},
		},
		{
			name: "Third",
			args: args{
				grid: [][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}},
			},
			want: []int{0, 1, 2, 3, 4, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindBall(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBall() = %v, want %v", got, tt.want)
			}
		})
	}
}
