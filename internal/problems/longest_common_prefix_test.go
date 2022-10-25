package problems

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				strings: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "2",
			args: args{
				strings: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			name: "3",
			args: args{
				strings: []string{"cir", "car"},
			},
			want: "c",
		},
		{
			name: "4",
			args: args{
				strings: []string{"reflower", "flow", "flight"},
			},
			want: "",
		},
		{
			name: "5",
			args: args{
				strings: []string{"flower", "flower", "flower", "flower"},
			},
			want: "flower",
		},
		{
			name: "6",
			args: args{
				strings: []string{"aa", "aa"},
			},
			want: "aa",
		},
		{
			name: "7",
			args: args{
				strings: []string{"ab", "a"},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.args.strings); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
