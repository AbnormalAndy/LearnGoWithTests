package main

import (
	"blackjack"
	"fmt"
	"strings"
)

func main() {
	greeting := `
	 ____  _     ____  ____  _  __    _  ____  ____  _  __  
	/  __\/ \   /  _ \/   _\/ |/ /   / |/  _ \/   _\/ |/ /  
	| | //| |   | / \||  /  |   /    | || / \||  /  |   /   
	| |_\\| |_/\| |-|||  \_ |   \ /\_| || |-|||  \_ |   \   
	\____/\____/\_/ \|\____/\_|\_\\____/\_/ \|\____/\_|\_\

	`

	fmt.Println(greeting)

	deck := [13]int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

	var dealerHand []int
	
	dealerHand = append(dealerHand, blackjack.DealCard(deck))
	// fmt.Println(dealerHand)

	dealerHand = append(dealerHand, blackjack.DealCard(deck))
	// fmt.Println(dealerHand)
	// TO-DO: Reveal only first number from the hand.
	fmt.Printf("DEALER hand: %d and X.\n\n", dealerHand[0])

	dealerSum := 0

	dealerSum = blackjack.AddCards(dealerHand)
	// fmt.Printf("Dealer's sum: %d.\n\n", dealerSum)

	for dealerSum < 17 {
		dealerHand = append(dealerHand, blackjack.DealCard(deck))
		// fmt.Printf("Dealer's hand: %v.\n\n", dealerHand)

		dealerSum = blackjack.AddCards(dealerHand)	
		// fmt.Printf("Dealer's sum: %d.\n\n", dealerSum)
	}


	var playerHand []int
	
	playerHand = append(playerHand, blackjack.DealCard(deck))
	// fmt.Println(playerHand)

	playerHand = append(playerHand, blackjack.DealCard(deck))
	// fmt.Println(playerHand)
	fmt.Printf("YOUR hand: %v.\n", playerHand)

	playerSum := 0

	playerSum = blackjack.AddCards(playerHand)
	// fmt.Printf("Player's sum: %d.\n\n", playerSum)

	finishHand := false

	// Asks if the player would like to continue playing.
	for finishHand != true {
		if playerSum > 21 {
			break
		}

		var playerDecision string
		
		fmt.Printf("\nYOUR card total is: %d.\n\n", playerSum)
		fmt.Println("Would you like to Hit ('H') or Stay ('S')?\n")

		fmt.Scanln(&playerDecision)

		playerDecision = strings.ToLower(playerDecision)

		if playerDecision == "hit" || playerDecision == "h" {
			playerHand = append(playerHand, blackjack.DealCard(deck))
			playerSum = blackjack.AddCards(playerHand)

			fmt.Printf("\nYOUR hand: %v.\n", playerHand)
			fmt.Printf("YOUR sum: %d.\n", playerSum)
		} else {
			finishHand = true
		}
	}


	if blackjack.EvaluateWinner(playerSum, dealerSum) {
		fmt.Printf("\nYou WIN!\n\n")
		fmt.Printf("YOUR hand: %v.\n", playerHand)
		fmt.Printf("YOUR sum: %d.\n\n", playerSum)
		fmt.Printf("DEALER hand: %v.\n", dealerHand) 
		fmt.Printf("DEALER sum: %d.\n\n", dealerSum)
	} else {
		fmt.Printf("\nYou LOSE!\n\n")
		fmt.Printf("YOUR hand: %v.\n", playerHand)
		fmt.Printf("YOUR sum: %d.\n\n", playerSum)
		fmt.Printf("DEALER hand: %v.\n", dealerHand) 
		fmt.Printf("DEALER sum: %d.\n\n", dealerSum)
	}

}
