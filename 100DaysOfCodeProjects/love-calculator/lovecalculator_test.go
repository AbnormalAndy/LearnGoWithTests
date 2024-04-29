package lovecalculator_test

import (
	"lovecalculator"
	"testing"
)

func TestAddLettersFromTRUEAndLOVE(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name1, name2 string
		want int
	}
	testCases := []testCase{
		{name1: "Brad Pitt", name2: "Jennifer Aniston", want: 73},
		{name1: "Prince William", name2: "Kate Middleton", want: 67},
		{name1: "Ashton Kutcher", name2: "Mila Kunis", want: 63},
		{name1: "Angela Yu", name2: "Jack Bauer", want: 53},
		{name1: "David Beckham", name2: "Victoria Beckham", want: 45},
		{name1: "Mario", name2: "Princess Peach", want: 43},
		{name1: "Kanye West", name2: "Kim Kardashian", want: 42},
		{name1: "Truth Truth", name2: "Love", want: 94},
		{name1: "QA", name2: "Violin", want: 3},
	}
	for _, tc := range testCases {
		got := lovecalculator.AddLetters(tc.name1, tc.name2)
		if tc.want != got {
			t.Errorf("AddLetters(%s, %s): Want %d, got %d.", tc.name1, tc.name2, tc.want, got)
		}
	}
}

func TestPrintStatementDependingOnScore(t *testing.T) {
	t.Parallel()
	type testCase struct {
		number int
		want string
	}
	testCases := []testCase{
		{number: 9, want: "Your score is 9, you go together like coke and mentos."},
		{number: 91, want: "Your score is 91, you go together like coke and mentos."},
		{number: 45, want: "Your score is 45, you are alright together."},
		{number: 39, want: "Your score is 39."},
		{number: 51, want: "Your score is 51."},
	}
	for _, tc := range testCases {
		got := lovecalculator.PrintStatement(tc.number)
		if tc.want != got {
			t.Errorf("PrintStatementDependingOnScore(%d): Want %s, got %s.", tc.number, tc.want, got)
		}
	}
}

