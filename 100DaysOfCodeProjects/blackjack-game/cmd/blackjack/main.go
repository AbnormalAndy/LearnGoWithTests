package main

import (
	"blackjack"
	"fmt"
	"strings"
)

func main() {
	deck := [13]int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

	var dealerHand []int
	
	dealerHand = append(dealerHand, blackjack.DealCard(deck))
	fmt.Println(dealerHand)

	dealerHand = append(dealerHand, blackjack.DealCard(deck))
	fmt.Println(dealerHand)
	// TO-DO: Reveal only first number from the hand.
	fmt.Printf("Dealer's hand: %d and X.\n\n", dealerHand[0])

	dealerSum := 0

	dealerSum = blackjack.AddCards(dealerHand)
	fmt.Printf("Dealer's sum: %d.\n\n", dealerSum)

	for dealerSum < 17 {
		dealerHand = append(dealerHand, blackjack.DealCard(deck))
		fmt.Printf("Dealer's hand: %v.\n\n", dealerHand)

		dealerSum = blackjack.AddCards(dealerHand)	
		fmt.Printf("Dealer's sum: %d.\n\n", dealerSum)
	}


	var playerHand []int
	
	playerHand = append(playerHand, blackjack.DealCard(deck))
	fmt.Println(playerHand)

	playerHand = append(playerHand, blackjack.DealCard(deck))
	fmt.Println(playerHand)
	fmt.Printf("Player's hand: %v.\n\n", playerHand)

	playerSum := 0

	playerSum = blackjack.AddCards(playerHand)
	fmt.Printf("Player's sum: %d.\n\n", playerSum)

	// TO-DO: Begin code to take user input here.
	// TO-DO: Evaluate "Hit / H" and "Stay / S" commands.
	finishHand := false

	for finishHand != true {
		if playerSum > 21 {
			break
		}

		var playerDecision string
		fmt.Printf("Your card total is: %d.\n\n", playerSum)
		fmt.Println("Would you like to Hit ('H') or Stay ('S')?\n")
		fmt.Scanln(&playerDecision)
		playerDecision = strings.ToLower(playerDecision)

		if playerDecision == "hit" || playerDecision == "h" {
			playerHand = append(playerHand, blackjack.DealCard(deck))
			playerSum = blackjack.AddCards(playerHand)
			fmt.Printf("Player's hand: %v.\n\n", playerHand)
			fmt.Printf("Player's sum: %d.\n\n", playerSum)
		} else {
			finishHand = true
		}
	}


	// This saves 8 lines of code.
	// BUG: If both hands are above 21, the player wins.
	if blackjack.EvaluateWinner(playerSum, dealerSum) {
		fmt.Printf("You win!\n")
		fmt.Printf("Player hand: %v. Player sum: %d.\n", playerHand, playerSum)
		fmt.Printf("Computer hand: %v. Dealer sum: %d.\n", dealerHand, dealerSum)
	} else {
		fmt.Printf("You lose!\n")
		fmt.Printf("Player hand: %v. Player sum: %d.\n", playerHand, playerSum)
		fmt.Printf("Computer hand: %v. Dealer sum: %d.\n", dealerHand, dealerSum)
	}

}
