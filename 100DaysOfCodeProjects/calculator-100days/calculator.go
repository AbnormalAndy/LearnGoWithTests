package calculator

import "errors"

func Add(x, y float64) float64 {
	return x + y
}

func Subtract(x, y float64) float64 {
	return x - y
}

func Multiply(x, y float64) float64 {
	return x * y
}

func Divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Division by zero not allowed.")
	}
	return x / y, nil
}
