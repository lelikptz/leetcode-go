package medium

import "testing"

func TestStockSpanner_Next(t *testing.T) {
	type args struct {
		items [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "First",
			args: args{
				items: [][]int{{100, 1}, {80, 1}, {60, 1}, {70, 2}, {60, 1}, {75, 4}, {85, 6}},
			},
		},
		{
			name: "Second",
			args: args{
				items: [][]int{{29, 1}, {91, 2}, {62, 1}, {76, 2}, {51, 1}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Constructor()
			for _, item := range tt.args.items {
				if val := s.Next(item[0]); val != item[1] {
					t.Errorf("TestStockSpanner() = %v, want %v", val, item[1])
				}
			}
		})
	}
}
