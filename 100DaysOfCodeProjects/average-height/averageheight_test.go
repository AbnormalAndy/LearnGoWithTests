package averageheight_test

import (
	"averageheight"
	"math"
	"testing"
)

func TestCalculateTotalHeight(t *testing.T) {
	t.Parallel()
	TotalHeight := []int{180, 124, 165, 173, 189, 169, 146}

	got := averageheight.CalculateTotalHeight(TotalHeight)
	want := 1146

	if want != got {
		t.Errorf("Want %d; got %d.", want, got)
	}
}

func TestCalculateNumberOfStudents(t *testing.T) {
	t.Parallel()
	TotalNumberOfStudents := []int{180, 124, 165, 173, 189, 169, 146}

	got := averageheight.CalculateNumberOfStudents(TotalNumberOfStudents)
	want := 7

	if want != got {
		t.Errorf("Want %d; got %d.", want, got)
	}
}

func TestCalculateAverage(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b	float64
		want	float64
	}
	testCases := []testCase{
		{a: 14.0, b: 2.0, want: 7.0},
		{a: 20.0, b: 5.0, want: 4.0},
		{a: 1146.0, b: 7.0, want: 163.714286},
	}
	for _, tc := range testCases {
		got, err := averageheight.CalculateAverage(tc.a, tc.b)
		if err != nil {
			t.Fatalf("Want no error for valid input, got %v.", err)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("CalculateAverage(%f, %f): Want %f; got %f.", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestInvalidCalculateAverage(t *testing.T) {
	t.Parallel()
	_, err := averageheight.CalculateAverage(1, 0)
	if err == nil {
		t.Error("Want error for invalid input, got nil")
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
