package gradingprogram

func GradeProgram(score int) string {
	switch {
	case score > 90:
		return "Outstanding"
	case score > 80:
		return "Exceeds Expectations"
	case score > 70:
		return "Acceptable"
	default:
		return "Fail"
	}
}
