package numberguessinggame_test

import (
	"numberguessinggame"
	"testing"
)

func TestEvaluateIfHigher(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b	int
		want	bool
	}
	testCases := []testCase{
		{a: 13, b: 11, want: false},
		{a: 13, b: 15, want: true},
		{a: 13, b: 13, want: false},
	}
	for _, tc := range testCases {
		got := numberguessinggame.EvaluateIfHigher(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Want %v; got %v.", tc.want, got)
		}
	}
}

func TestEvaluateIfLower(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b	int
		want	bool
	}
	testCases := []testCase{
		{a: 13, b: 11, want: true},
		{a: 13, b: 13, want: false},
		{a: 13, b: 15, want: false},
	}
	for _, tc := range testCases {
		got := numberguessinggame.EvaluateIfLower(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Want %v; got %v.", tc.want, got)
		}
	}
}

func TestEvaluateIfEqual(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b	int
		want	bool
	}
	testCases := []testCase{
		{a: 13, b: 11, want: false},
		{a: 13, b: 13, want: true},
		{a: 13, b: 15, want: false},
	}
	for _, tc := range testCases {
		got := numberguessinggame.EvaluateIfEqual(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Want %v; got %v.", tc.want, got)
		}
	}
}

