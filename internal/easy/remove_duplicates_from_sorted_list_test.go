package easy

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	type args struct {
		head *structure.ListNode
	}
	tests := []struct {
		name string
		args args
		want *structure.ListNode
	}{
		{
			name: "First",
			args: args{
				head: structure.CreateList(1, 1, 2),
			},
			want: structure.CreateList(1, 2),
		},
		{
			name: "Second",
			args: args{
				head: structure.CreateList(1, 1, 2, 3, 3),
			},
			want: structure.CreateList(1, 2, 3),
		},
		{
			name: "Third",
			args: args{
				head: structure.CreateList(1, 1, 1),
			},
			want: structure.CreateList(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
