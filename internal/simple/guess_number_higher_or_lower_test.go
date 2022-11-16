package simple

import (
	"testing"
)

func TestGuessNumber(t *testing.T) {
	type args struct {
		n int
	}
	test := struct {
		name string
		args args
		want int
	}{
		name: "Third",
		args: args{
			n: 1,
		},
		want: 10241024,
	}
	GuessValue = 10241024

	t.Run("First", func(t *testing.T) {
		if got := GuessNumber(test.args.n); got != test.want {
			t.Errorf("GuessNumber() = %v, want %v", got, test.want)
		}
	})
}
