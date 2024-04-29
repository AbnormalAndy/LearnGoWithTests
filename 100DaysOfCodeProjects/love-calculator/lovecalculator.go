package lovecalculator

import (
	"fmt"
	"strings"
)

func AddLetters(name1, name2 string) int {
	lowerName1 := strings.ToLower(name1)
	lowerName2 := strings.ToLower(name2)

	t := strings.Count(lowerName1, "t") + strings.Count(lowerName2, "t")
	r := strings.Count(lowerName1, "r") + strings.Count(lowerName2, "r")
	u := strings.Count(lowerName1, "u") + strings.Count(lowerName2, "u")
	e1 := strings.Count(lowerName1, "e") + strings.Count(lowerName2, "e")

	l := strings.Count(lowerName1, "l") + strings.Count(lowerName2, "l")
	o := strings.Count(lowerName1, "o") + strings.Count(lowerName2, "o")
	v := strings.Count(lowerName1, "v") + strings.Count(lowerName2, "v")
	e2 := strings.Count(lowerName1, "e") + strings.Count(lowerName2, "e")

	firstNumber := (t + r + u + e1) * 10
	secondNumber := l + o + v + e2

	return firstNumber + secondNumber
}

func PrintStatement(number int) string {
	if number < 10 || number > 90 {
		return fmt.Sprintf("Your score is %d, you go together like coke and mentos.\n\n", number)
	} else if number > 40 && number < 50 {
		return fmt.Sprintf("Your score is %d, you are alright together.\n\n", number)
	} else {
		return fmt.Sprintf("Your score is %d.\n\n", number)
	}
}

