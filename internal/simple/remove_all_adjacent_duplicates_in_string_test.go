package simple

import "testing"

func TestRemoveAllDuplicatesInString(t *testing.T) {
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
				s: "abbaca",
			},
			want: "ca",
		},
		{
			name: "Second",
			args: args{
				s: "azxxzy",
			},
			want: "ay",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAllDuplicatesInString(tt.args.s); got != tt.want {
				t.Errorf("RemoveAllDuplicatesInString() = %v, want %v", got, tt.want)
			}
		})
	}
}
