package medium

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "First",
			args: args{
				s: []string{""},
			},
			want: [][]string{{""}},
		},
		{
			name: "Second",
			args: args{
				s: []string{"a"},
			},
			want: [][]string{{"a"}},
		},
		{
			name: "Third",
			args: args{
				s: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"eat", "ate", "tea"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.args.s)
			sort.Slice(got, func(i, j int) bool {
				sort.Strings(got[i])
				sort.Strings(got[j])
				return len(got[i]) < len(got[j])
			})
			sort.Slice(tt.want, func(i, j int) bool {
				sort.Strings(tt.want[i])
				sort.Strings(tt.want[j])
				return len(got[i]) < len(got[j])
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
