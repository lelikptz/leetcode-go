package easy

import "testing"

func TestClimbStairs(t *testing.T) {
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
				n: 2,
			},
			want: 2,
		},
		{
			name: "Second",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "Third",
			args: args{
				n: 38,
			},
			want: 63245986,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairs(tt.args.n); got != tt.want {
				t.Errorf("ClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
