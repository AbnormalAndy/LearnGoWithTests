package main

import (
	"fmt"
	"headsortails"
	"math/rand"
)

func main() {
	RandomNumber := rand.Intn(2)

	fmt.Printf(headsortails.CallsHeadsOrTails(RandomNumber))
}

