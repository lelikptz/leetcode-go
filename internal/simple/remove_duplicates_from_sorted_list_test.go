package simple

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "First",
			args: args{
				head: CreateList(1, 1, 2),
			},
			want: CreateList(1, 2),
		},
		{
			name: "Second",
			args: args{
				head: CreateList(1, 1, 2, 3, 3),
			},
			want: CreateList(1, 2, 3),
		},
		{
			name: "Third",
			args: args{
				head: CreateList(1, 1, 1),
			},
			want: CreateList(1),
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
