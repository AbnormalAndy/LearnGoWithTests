package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TipCalculator(totalBill float64, desiredTip float64, numberOfPeople float64) float64 {
	totalTip := (desiredTip / 100) + 1
	return (totalBill / numberOfPeople) * totalTip
}

func main() {
	var tip float64
	var people float64

	fmt.Print("Welcome to the tip calculator!\n\n")

	// Collects total bill in a float value.
	fmt.Println("What was the total bill?")
	inputReaderOne := bufio.NewReader(os.Stdin)
	billInput, err := inputReaderOne.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	billInput = strings.TrimSuffix(billInput, "\n")
	bill, err := strconv.ParseFloat(billInput, 64)
	if err != nil {
		fmt.Println(err)
	}

	// Collects a whole number for the tip.
	fmt.Println("How much tip would you like to give? 10, 12, 15, 20?")
	fmt.Scanln(&tip)

	// Collects the number of people to split the bill between.
	fmt.Println("How many people to split the bill?")
	fmt.Scanln(&people)

	// fmt.Printf("Bill: %.2f - Tip %.2f - Peolpe %.2f\n", bill, tip, people)

	fmt.Printf("Each Person Should Pay: $%.2f\n", TipCalculator(bill, tip, people))
}
