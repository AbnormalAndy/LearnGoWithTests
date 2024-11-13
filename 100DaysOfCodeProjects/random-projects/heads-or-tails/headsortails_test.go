package headsortails_test

import (
	"headsortails"
	"testing"
)

func TestCallsHeadsOrTails(t *testing.T) {
	t.Parallel()
	type testCase struct {
		number int
		want string
	}
	testCases := []testCase{
		{number: 0, want: "Heads\n"},
		{number: 1, want: "Tails\n"},
        	{number: -1, want: "Number not 0 or 1.\n"},
		{number: 2, want: "Number not 0 or 1.\n"},
	}
	for _, tc := range testCases {
		got := headsortails.CallsHeadsOrTails(tc.number)
		if tc.want != got {
			t.Errorf("CallsHeadsOrTails(%d): Want %s; got %s.", tc.number, tc.want, got)
		}
	}
}

