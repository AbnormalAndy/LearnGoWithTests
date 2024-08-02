package blackjack_test

import (
	"blackjack"
	"testing"
)

func TestAddCards(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b	int
		want	int
	}
	testCases := []testCase{
		{a: 11, b: 2, want: 13},
		{a: 5, b: 7, want: 12},
		{a: 10, b: 11, want: 21},
		{a: 11, b: 13, want: 14},
		{a: 15, b: 11, want: 16},
		{a: 11, b: 11, want: 12},
	}
	for _, tc := range testCases {
		got := blackjack.AddCards(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Want: %d; got: %d.", tc.want, got)
		}
	}
}

