package gradingprogram_test

import (
	"gradingprogram"
	"testing"
)

func TestScores(t *testing.T) {
	t.Parallel()
	_ = gradingprogram.Scores{
		Name:	"Hermione",
		Score:	99,
	}
}

func TestScoreToGrade(t *testing.T) {
	t.Parallel()
	g := gradingprogram.Scores{
		Name:	"Harry",
		Score:	81,
	}
	want := "Exceeds Expectations"
	got := gradingprogram.ScoreToGrade(g)
	if want != got {
		t.Errorf("Want %s; got %s.", want, got)
	}
}
