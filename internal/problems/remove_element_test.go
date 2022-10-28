package problems

import "testing"

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums         []int
		expectedNums []int
		val          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First",
			args: args{
				nums:         []int{3, 2, 2, 3},
				expectedNums: []int{2, 2},
				val:          3,
			},
			want: 2,
		},
		{
			name: "Second",
			args: args{
				nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
				expectedNums: []int{0, 1, 4, 0, 3},
				val:          3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveElement(tt.args.nums, tt.args.val)
			if got != tt.want {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}

			for i := 0; i < got; i++ {
				if tt.args.nums[i] != tt.args.expectedNums[i] {
					t.Errorf("error element = %v, want %v", tt.args.nums[i], tt.args.expectedNums[i])
				}
			}
		})
	}
}
