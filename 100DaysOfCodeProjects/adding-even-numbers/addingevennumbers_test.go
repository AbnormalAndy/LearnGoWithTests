package addingevennumbers_test

import (
	"addingevennumbers"
	"testing"
)

func TestAddEvenNumbers(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a	int
		want	int
	}
	testCases := []testCase{
		{a: 10, want: 30},
		{a: 52, want: 702},
	}
	for _, tc := range testCases {
		got := addingevennumbers.AddEvenNumbers(tc.a)
		if tc.want != got {
			t.Errorf("AddEvenNumbers(%d): Want %d; got %d.", tc.a, tc.want, got)
		}
	}
}
