package gradingprogram_test

import (
	"gradingprogram"
	"testing"
)

func TestScores(t *testing.T) {
	t.Parallel()
	_ = gradingprogram.Student{
		Name:	"Hermione",
		Score:	99,
	}
}

// To-Do 1: Intake a struct and output a struct.
// To-Do 2: Make a test case that tests all the ranges.
func TestScoreToGrade(t *testing.T) {
	t.Parallel()
	g := gradingprogram.Student{
		Name:	"Harry",
		Score:	81,
		Grade: "Exceeds Expectations",
	}
	want := "Exceeds Expectations"
	result := gradingprogram.ScoreToGrade(g)
	got := result.Grade
	if want != got {
		t.Errorf("Want %s; got %s.", want, got)
	}
}
