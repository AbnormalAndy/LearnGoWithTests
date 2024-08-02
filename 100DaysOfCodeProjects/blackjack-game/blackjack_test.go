package blackjack_test

import (
	"blackjack"
	"testing"
)

func TestAddCards(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a	[]int
		want	int
	}
	testCases := []testCase{
		{a: []int{11, 2}, want: 13},
		{a: []int{5, 7}, want: 12},
		{a: []int{10, 11}, want: 21},
		{a: []int{11, 13}, want: 14},
		{a: []int{15, 11}, want: 16},
		{a: []int{11, 11}, want: 12},
	}
	for _, tc := range testCases {
		got := blackjack.AddCards(tc.a)
		if tc.want != got {
			t.Errorf("Want: %d; got: %d.", tc.want, got)
		}
	}
}

func TestDealCard(t *testing.T) {
	t.Parallel()
	deck := [13]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	want := 1
	got := blackjack.DealCard(deck)
	if want != got {
		t.Errorf("Want: %d; got: %d.", want, got)
	}
}
