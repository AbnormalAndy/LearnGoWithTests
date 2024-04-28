package main

import (
	"bandname"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var city string
	var petname string

	fmt.Print("BAND NAME GENERATOR v0.2\n\n")

	fmt.Print("What city did you grow up in?\n")
	inputReaderOne := bufio.NewReader(os.Stdin)
	city, _ = inputReaderOne.ReadString('\n')

	fmt.Print("What is your pet's name?\n")
	inputReaderTwo := bufio.NewReader(os.Stdin)
	petname, _ = inputReaderTwo.ReadString('\n')

	city = strings.TrimSuffix(city, "\n")
	petname = strings.TrimSuffix(petname, "\n")

	fmt.Printf("\nYou band name is %s.\n", bandname.CombiningCityAndPetName(city, petname))

}
