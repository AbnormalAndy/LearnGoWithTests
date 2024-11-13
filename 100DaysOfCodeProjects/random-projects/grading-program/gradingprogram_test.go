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

// Intake an array of students and output new array with student grades.
