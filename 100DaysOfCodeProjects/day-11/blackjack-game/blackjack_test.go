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
		{a: []int{11, 11, 9}, want: 21},
		{a: []int{11, 11, 11}, want: 13},
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

func TestEvaluateWinner(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b	int
		want	bool
	}
	testCases := []testCase{
		{a: 21, b: 21, want: true},
		{a: 20, b: 21, want: false},
		{a: 20, b: 19, want: true},
		{a: 19, b: 20, want: false},
		{a: 17, b: 25, want: true},
		{a: 24, b: 17, want: false},
		{a: 26, b: 24, want: false},
	}
	for _, tc := range testCases {
		got := blackjack.EvaluateWinner(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Player sum: %d; Dealer sum: %d; want: %v; got: %v.", tc.a, tc.b, tc.want, got)
		}
	}
}

