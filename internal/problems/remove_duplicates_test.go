package problems

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums         []int
		expectedNums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First",
			args: args{
				nums:         []int{1, 1, 2},
				expectedNums: []int{1, 2},
			},
			want: 2,
		},
		{
			name: "First",
			args: args{
				nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
				expectedNums: []int{0, 1, 2, 3, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicates(tt.args.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}

			for i := 0; i < got; i++ {
				if tt.args.nums[i] != tt.args.expectedNums[i] {
					t.Errorf("error element = %v, want %v", tt.args.nums[i], tt.args.expectedNums[i])
				}
			}
		})
	}
}
