package medium

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "Second",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "Third",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "Fourth",
			args: args{
				s: "aabaab!bb",
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
