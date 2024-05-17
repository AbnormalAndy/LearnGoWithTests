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
		got, err := addingevennumbers.AddEvenNumbers(tc.a)
		if err != nil {
			t.Fatalf("Want no error for valid input; got %v.", err)
		}
		if tc.want != got {
			t.Errorf("AddEvenNumbers(%d): Want %d; got %d.", tc.a, tc.want, got)
		}
	}
}

func TestAddEvenNumbersInvalid(t *testing.T) {
	t.Parallel()
	_, err := addingevennumbers.AddEvenNumbers(-1)
	if err == nil {
		t.Error("Want error for invalid input; got nil.")
	}
}

func TestAddEvenNumbersMaxInvalid(t *testing.T) {
	t.Parallel()
	_, err := addingevennumbers.AddEvenNumbers(1001)
	if err == nil {
		t.Error("Want error for invalid input; got nil.")
	}
}
