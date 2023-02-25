package medium

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "First",
			s:    "",
			want: "",
		},
		{
			name: "Second",
			s:    "a",
			want: "a",
		},
		{
			name: "Third",
			s:    "ab",
			want: "a",
		},
		{
			name: "Fourth",
			s:    "aa",
			want: "aa",
		},
		{
			name: "Fifth",
			s:    "abcba",
			want: "abcba",
		},
		{
			name: "Sixth",
			s:    "zabcdcb",
			want: "bcdcb",
		},
		{
			name: "Seventh",
			s:    "bcdcbza",
			want: "bcdcb",
		},
		{
			name: "Eighth",
			s:    "bcdcbzaba",
			want: "bcdcb",
		},
		{
			name: "Ninth",
			s:    "brcccccbran",
			want: "ccccc",
		},
		{
			name: "Tenth",
			s:    "abaffggffhah",
			want: "ffggff",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome(tt.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
