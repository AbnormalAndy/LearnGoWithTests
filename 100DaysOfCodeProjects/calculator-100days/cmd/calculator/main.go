package main

import (
	"calculator"
	"fmt"
)

func main() {

	x := 3.0
	y := 5.0
	input := "+"
	total := 0.0

	switch {
	case input == "+":
		total += calculator.Add(x, y)
	case input == "-":
		calculator.Subtract(x, y)
	case input == "*":
		calculator.Multiply(x, y)
	case input == "/":
		calculator.Divide(x, y)
	}

	fmt.Printf("Total: %f", total)
}
