package gradingprogram

type Scores struct {
	Name	string
	Score	int
	Grade	string
}

// Intake an array of structs and out put the same array.
func ScoreToGrade(s Scores) string {
	switch {
	case s.Score > 90:
		return "Outstanding"
	case s.Score > 80:
		return "Exceeds Expectations"
	case s.Score > 70:
		return "Acceptable"
	default:
		return "Fail"
	}
}
