package simple

import (
	"github.com/lelikptz/leetcode-go/internal/structure"
	"testing"
)

func TestHasCycle(t *testing.T) {
	type args struct {
		head *structure.ListNode
	}

	a := &structure.ListNode{Val: 1}
	b := &structure.ListNode{Val: 1, Next: a}
	a.Next = b

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "First",
			args: args{
				head: &structure.ListNode{Val: 1, Next: &structure.ListNode{Val: 1, Next: &structure.ListNode{Val: 1}}},
			},
			want: false,
		},
		{
			name: "Second",
			args: args{
				head: b,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycle(tt.args.head); got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
