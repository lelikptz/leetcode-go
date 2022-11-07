package simple

import "testing"

func TestMaximum69Number(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First",
			args: args{
				num: 9669,
			},
			want: 9969,
		},
		{
			name: "Second",
			args: args{
				num: 9996,
			},
			want: 9999,
		},
		{
			name: "Third",
			args: args{
				num: 9999,
			},
			want: 9999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Maximum69Number(tt.args.num); got != tt.want {
				t.Errorf("Maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
