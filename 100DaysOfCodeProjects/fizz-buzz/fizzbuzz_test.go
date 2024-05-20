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
		{a: 6, want: "Fizz"},
		{a: 9, want: "Fizz"},
		{a: 10, want: "Buzz"},
		{a: 12, want: "Fizz"},
		{a: 15, want: "FizzBuzz"},
	}
	for _, tc := range testCases {
		got, err := fizzbuzz.FizzBuzz(tc.a)
		if err != nil {
			t.Fatalf("FizzBuzz(%d): Want no error for valid input; got %v.", tc.a, got)
		}
		if tc.want != got {
			t.Errorf("FizzBuzz(%d): Want %s; got %s.", tc.a, tc.want, got)
		}
	}
}

func TestFizzBuzzInvalidInput(t *testing.T) {
	t.Parallel()
	_, err := fizzbuzz.FizzBuzz(7)
	if err == nil {
		t.Error("Want error for invalid input; got nil.")
	}
}
