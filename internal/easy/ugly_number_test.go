package easy

import "testing"

func TestIsUgly(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "First",
			args: args{
				n: 6,
			},
			want: true,
		},
		{
			name: "Second",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "Third",
			args: args{
				n: 14,
			},
			want: false,
		},
		{
			name: "Forth",
			args: args{
				n: 8,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUgly(tt.args.n); got != tt.want {
				t.Errorf("IsUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
