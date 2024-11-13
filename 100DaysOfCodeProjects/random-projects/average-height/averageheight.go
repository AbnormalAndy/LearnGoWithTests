package averageheight

import (
	"errors"
)

func CalculateTotalHeight(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum
}

func CalculateNumberOfStudents(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += 1
	}
	return sum
}

func CalculateAverage(NumberOne, NumberTwo float64) (float64, error) {
	if NumberTwo == 0 {
		return 0, errors.New("Division by zero not allowed.")
	}

	return NumberOne / NumberTwo, nil
}
