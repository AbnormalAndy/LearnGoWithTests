package gradingprogram_test

import (
	"gradingprogram"
	"testing"
)

func TestGradeProgram(t *testing.T) {
	t.Parallel()
	type testCase struct {
		score	int
		want	string
	}
	testCases := []testCase{
		{score: 92, want: "Outstanding"},
		{score: 82, want: "Exceeds Expectations"},
		{score: 72, want: "Acceptable"},
		{score: 62, want: "Fail"},
	}
	for _, tc := range testCases {
		got := gradingprogram.GradeProgram(tc.score)
		if tc.want != got {
			t.Errorf("Score %d; Want %s; Got %s.", tc.score, tc.want, got)
		}
	}
}
