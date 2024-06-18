package gradingprogram

type Student struct {
	Name	string
	Score	int
	Grade	string
}

// Intake an array of structs and out put the same array.
func ScoreToGrade(s Student) Student {
	switch {
	case s.Score > 90:
		s.Grade = "Outstanding"
		return s 
	case s.Score > 80:
		s.Grade = "Exceeds Expectations"
		return s
	case s.Score > 70:
		s.Grade = "Acceptable"
		return s
	default:
		s.Grade = "Fail"
		return s
	}
}
