package medium

import "testing"

func TestMinMutation(t *testing.T) {
	type args struct {
		startGene string
		endGene   string
		bank      []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First",
			args: args{
				startGene: "AACCGGTT",
				endGene:   "AACCGGTA",
				bank:      []string{"AACCGGTA"},
			},
			want: 1,
		},
		{
			name: "Second",
			args: args{
				startGene: "AACCGGTT",
				endGene:   "AAACGGTA",
				bank:      []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			},
			want: 2,
		},
		{
			name: "Third",
			args: args{
				startGene: "AAAAAAAA",
				endGene:   "AAAAACGG",
				bank:      []string{"AAAAAAGA", "AAAAAGGA", "AAAAACGA", "AAAAACGG", "AAAAAAGG", "AAAAAAGC"},
			},
			want: 3,
		},
		{
			name: "Forth",
			args: args{
				startGene: "AAAAAAAA",
				endGene:   "CCCCCCCC",
				bank:      []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC", "AAAACCCC", "AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA"},
			},
			want: -1,
		},
		{
			name: "Five",
			args: args{
				startGene: "AAAACCCC",
				endGene:   "CCCCCCCC",
				bank:      []string{"AAAACCCA", "AAACCCCA", "AACCCCCA", "AACCCCCC", "ACCCCCCC", "CCCCCCCC", "AAACCCCC", "AACCCCCC"},
			},
			want: 4,
		},
		{
			name: "Six",
			args: args{
				startGene: "AACCGGTT",
				endGene:   "AAACGGTA",
				bank:      []string{"AACCGATT", "AACCGATA", "AAACGATA", "AAACGGTA"},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMutation(tt.args.startGene, tt.args.endGene, tt.args.bank); got != tt.want {
				t.Errorf("MinMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
