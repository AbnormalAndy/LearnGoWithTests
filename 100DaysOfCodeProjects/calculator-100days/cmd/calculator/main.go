package main

import (
	"bufio"
	"calculator"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Ask for the first number.
	fmt.Println("What is the first number?\n")
	inputReaderOne := bufio.NewReader(os.Stdin)
	firstNumberInput, err := inputReaderOne.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	firstNumberInput = strings.TrimSuffix(firstNumberInput, "\n")
	firstNumber, err := strconv.ParseFloat(firstNumberInput, 64)
	if err != nil {
		fmt.Println(err)
	}

	
	sum := 0.0

	quit := false

	for quit != true {

		// Ask for the second number.
		fmt.Println("What is the second number?\n")
		inputReaderTwo := bufio.NewReader(os.Stdin)
		secondNumberInput, err := inputReaderTwo.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		secondNumberInput = strings.TrimSuffix(secondNumberInput, "\n")
		secondNumber, err := strconv.ParseFloat(secondNumberInput, 64)
		if err != nil {
			fmt.Println(err)
		}

		var operator string 

		fmt.Println("Choose an Operator: + - * /\n")
		fmt.Scanln(&operator)


		switch {
		case operator == "+":
			sum = calculator.Add(firstNumber, secondNumber)
		case operator == "-":
			sum = calculator.Subtract(firstNumber, secondNumber)
		case operator == "*":
			sum = calculator.Multiply(firstNumber, secondNumber)
		case operator == "/":
			sum, _ = calculator.Divide(firstNumber, secondNumber)
		}
		fmt.Printf("Sum: %.2f.\n", sum)

		
		var calculateMore string

		fmt.Println("Would you like to continue? Yes or No\n")
		fmt.Scanln(&calculateMore)

		if strings.ToLower(calculateMore) == "no" || strings.ToLower(calculateMore) == "n" {
			quit = true
		}

		if strings.ToLower(calculateMore) != "yes" || strings.ToLower(calculateMore) == "y" {
			fmt.Println("Did not enter 'Yes' or 'No'.")
			quit = true
		}

		firstNumber = sum
	}
}
