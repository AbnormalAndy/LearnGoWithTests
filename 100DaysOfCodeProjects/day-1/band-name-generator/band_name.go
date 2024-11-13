package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var city string
	var pet string

	fmt.Print("BAND NAME GENERATOR v0.1\n\n")

	// Bufio allows for a multiple word input.
	fmt.Print("What city did you grow up in?\n")
	inputReaderOne := bufio.NewReader(os.Stdin)
	city, _ = inputReaderOne.ReadString('\n')

	fmt.Print("\nWhat is the name of your pet?\n")
	inputReaderTwo := bufio.NewReader(os.Stdin)
	pet, _ = inputReaderTwo.ReadString('\n')

	// Removes the newline from the user's answers.
	city = strings.TrimSuffix(city, "\n")
	pet = strings.TrimSuffix(pet, "\n")

	fmt.Printf("\nYour band name is %s %s.\n\n", city, pet)

}
