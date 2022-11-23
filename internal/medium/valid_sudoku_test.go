package medium

import "testing"

func TestIsValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "First",
			args: args{
				board: [][]byte{
					{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
					{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
					{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
					{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
					{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
					{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
					{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
					{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
					{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
				},
			},
			want: true,
		},
		{
			name: "Second",
			args: args{
				board: [][]byte{
					{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
					{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
					{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
					{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
					{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
					{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
					{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
					{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
					{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
				},
			},
			want: false,
		},
		{
			name: "Third",
			args: args{
				board: [][]byte{
					{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
					{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
					{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
					{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
					{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
					{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
					{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
				},
			},
			want: false,
		},
		{
			name: "Fourth",
			args: args{
				board: [][]byte{
					{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
					{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
					{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
					{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("IsValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
