package headsortails

func CallsHeadsOrTails(number int) string {
	switch {
	case number == 0:
		return "Heads\n"
	case number == 1:
		return "Tails\n"
	default:
		return "Number not 0 or 1.\n"
	}
}
