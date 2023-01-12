package medium

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *structure.ListNode
		l2 *structure.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structure.ListNode
	}{
		{
			name: "First",
			args: args{
				l1: structure.CreateList(2, 4, 3),
				l2: structure.CreateList(5, 6, 4),
			},
			want: structure.CreateList(7, 0, 8),
		},
		{
			name: "Second",
			args: args{
				l1: structure.CreateList(0),
				l2: structure.CreateList(0),
			},
			want: structure.CreateList(0),
		},
		{
			name: "Third",
			args: args{
				l1: structure.CreateList(9, 9, 9, 9, 9, 9, 9),
				l2: structure.CreateList(9, 9, 9, 9),
			},
			want: structure.CreateList(8, 9, 9, 9, 0, 0, 0, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
