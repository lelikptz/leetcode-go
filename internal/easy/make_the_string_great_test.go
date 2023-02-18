package easy

import "testing"

func TestMakeGood(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "First",
			args: args{
				s: "leEeetcode",
			},
			want: "leetcode",
		},
		{
			name: "First",
			args: args{
				s: "abBAcC",
			},
			want: "",
		},
		{
			name: "First",
			args: args{
				s: "s",
			},
			want: "s",
		},
		{
			name: "First",
			args: args{
				s: "Pp",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeGood(tt.args.s); got != tt.want {
				t.Errorf("MakeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
