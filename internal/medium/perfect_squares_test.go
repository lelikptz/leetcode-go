package medium

import "testing"

func TestNumSquares(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First",
			args: args{
				n: 12,
			},
			want: 3,
		},
		{
			name: "Second",
			args: args{
				n: 13,
			},
			want: 2,
		},
		{
			name: "Third",
			args: args{
				n: 43,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumSquares(tt.args.n); got != tt.want {
				t.Errorf("NumSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
