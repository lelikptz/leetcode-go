package easy

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		a *structure.ListNode
		b *structure.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structure.ListNode
	}{
		{
			name: "First",
			args: args{
				a: structure.CreateList(1, 2, 4),
				b: structure.CreateList(1, 3, 4),
			},
			want: structure.CreateList(1, 1, 2, 3, 4, 4),
		},
		{
			name: "Second",
			args: args{
				a: structure.CreateList(1, 2, 3, 4),
				b: structure.CreateList(5, 6, 7),
			},
			want: structure.CreateList(1, 2, 3, 4, 5, 6, 7),
		},
		{
			name: "Third",
			args: args{
				a: structure.CreateList(7, 82),
				b: structure.CreateList(5, 32, 79),
			},
			want: structure.CreateList(5, 7, 32, 79, 82),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTwoLists(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
