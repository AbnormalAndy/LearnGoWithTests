package fizzbuzz_test

import (
	"fizzbuzz"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a	int
		want	string
	}
	testCases := []testCase{
		{a: 3, want: "Fizz"},
		{a: 5, want: "Buzz"},
		{a: 15, want: "FizzBuzz"},
	}
	for _, tc := range testCases {
		got := fizzbuzz.FizzBuzz(tc.a)
		if tc.want != got {
			t.Errorf("FizzBuzz(%d): Want %s; got %s.", tc.a, tc.want, got)
		}
	}
}
