package main

import (
	"addingevennumbers"
	"fmt"
)

func main() {
	var numberFromUser int

	fmt.Print("\nAdd Even Numbers\n\n")

	fmt.Print("Enter a Number (0 - 1000): ")
	fmt.Scanln(&numberFromUser)

	sum, err := addingevennumbers.AddEvenNumbers(numberFromUser)
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n\nThe sum of even numbers for %d is %d.\n\n", numberFromUser, sum)
	}
}

