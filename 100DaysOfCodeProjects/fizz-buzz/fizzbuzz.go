package fizzbuzz

import (
	"errors"
)

func FizzBuzz(number int) (string, error) {
	switch {
	case number % 3 == 0 && number % 5 == 0:
		return "FizzBuzz", nil
	case number % 3 == 0:
		return "Fizz", nil
	case number % 5 == 0:
		return "Buzz", nil
	}
	return "", errors.New("Number is not cleanly divisble by 3 and/or 5.")
}
