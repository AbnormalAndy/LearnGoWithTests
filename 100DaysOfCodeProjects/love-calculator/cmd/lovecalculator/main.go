package main

import (
	"bufio"
	"fmt"
	"lovecalculator"
	"os"
	"strings"
)

func main() {
	var name1 string
	var name2 string

	fmt.Print("\nWelcome to the LOVE Calculator!\n\n")

	fmt.Print("Enter the first name.\n")
	inputReaderOne := bufio.NewReader(os.Stdin)
	name1, _ = inputReaderOne.ReadString('\n')

	fmt.Print("Enter the second name.\n")
	inputReaderTwo := bufio.NewReader(os.Stdin)
	name2, _ = inputReaderTwo.ReadString('\n')

	name1 = strings.TrimSuffix(name1, "\n")
	name2 = strings.TrimSuffix(name2, "\n")

	totalofletters := lovecalculator.AddLetters(name1, name2)

	fmt.Print(lovecalculator.PrintStatement(totalofletters))
}

