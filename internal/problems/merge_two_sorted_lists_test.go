package problems

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		a *ListNode
		b *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "First",
			args: args{
				a: createList(1, 2, 4),
				b: createList(1, 3, 4),
			},
			want: createList(1, 1, 2, 3, 4, 4),
		},
		{
			name: "Second",
			args: args{
				a: createList(1, 2, 3, 4),
				b: createList(5, 6, 7),
			},
			want: createList(1, 2, 3, 4, 5, 6, 7),
		},
		{
			name: "Third",
			args: args{
				a: createList(7, 82),
				b: createList(5, 32, 79),
			},
			want: createList(5, 7, 32, 79, 82),
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

func createList(elements ...int) *ListNode {
	slice := make([]ListNode, len(elements), len(elements))

	for i, el := range elements {
		slice[i] = ListNode{
			Val: el,
		}
	}

	for index := range slice {
		if len(slice) > index+1 {
			slice[index].Next = &slice[index+1]
		}
	}

	return &slice[0]
}
