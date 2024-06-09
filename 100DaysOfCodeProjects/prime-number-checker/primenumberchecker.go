package primenumberchecker

import (
	"math"
)

func PrimeNumberChecker(number int) string {
	if number < 2 {
		return "Number must be greater than 2."
	}
	sq_root := int(math.Sqrt(float64(number)))
	for i := 2; i <= sq_root; i++ {
		if number % i == 0 {
			return "Not a prime number."
		}
	}
	return "It is a prime number."
}
