package simple

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "First",
			args: args{
				rowIndex: 3,
			},
			want: []int{1, 3, 3, 1},
		},
		{
			name: "Second",
			args: args{
				rowIndex: 1,
			},
			want: []int{1, 1},
		},
		{
			name: "Third",
			args: args{
				rowIndex: 0,
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
