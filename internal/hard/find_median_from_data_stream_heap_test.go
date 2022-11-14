package hard

import "testing"

func TestMedianFinderHeap_FindMedian(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "First",
			args: args{
				nums: []int{1, 2},
			},
			want: 1.5,
		},
		{
			name: "Second",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "Third",
			args: args{
				nums: []int{6, 10, 2, 6, 5, 0},
			},
			want: 5.5,
		},
		{
			name: "Fourth",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: 5.5,
		},
		{
			name: "Five",
			args: args{
				nums: []int{12, 10, 13, 11, 5, 15, 1, 11, 6, 17, 14, 8, 17, 6, 4, 16, 8, 10, 2, 12},
			},
			want: 10.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMedianFinderHeap()
			for _, val := range tt.args.nums {
				m.AddNum(val)
			}

			if got := m.FindMedian(); got != tt.want {
				t.Errorf("FindMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
