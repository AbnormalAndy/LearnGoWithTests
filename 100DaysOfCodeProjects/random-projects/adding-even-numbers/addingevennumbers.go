package addingevennumbers

import (
	"errors"
)

func AddEvenNumbers(number int) (int, error) {
	sum := 0
	switch {
	case number < 0:
		return 0, errors.New("A negative number is not allowed.")
	case number > 1000:
		return 0, errors.New("Invalid number. Max number is 1000.")
	default:
		for i := range (number + 1) {
			if i % 2 == 0 {
				sum += i
			}
		}
		return sum, nil
	}
}
