package main

import (
	"blackjack"
	"fmt"
)

func main() {
	deck := [13]int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

	var dealerHand []int
	
	dealerHand = append(dealerHand, blackjack.DealCard(deck))
	fmt.Println(dealerHand)

	dealerHand = append(dealerHand, blackjack.DealCard(deck))
	fmt.Println(dealerHand)
	fmt.Printf("Dealer's hand: %v.\n\n", dealerHand)

	dealerSum := 0

	dealerSum = blackjack.AddCards(dealerHand)
	fmt.Printf("Dealer's sum: %d.\n\n", dealerSum)

	for dealerSum < 17 {
		dealerHand = append(dealerHand, blackjack.DealCard(deck))
		fmt.Printf("Dealer's hand: %v.\n\n", dealerHand)

		dealerSum = blackjack.AddCards(dealerHand)	
		fmt.Printf("Dealer's sum: %d.\n\n", dealerSum)
	}

	// TO-DO: Code the player hand and input.
	/*
	playerCardOne := blackjack.DealCard(deck)
	playerCardTwo := blackjack.DealCard(deck)
	fmt.Printf("Player's card one: %d; card two: %d.\n", playerCardOne, playerCardTwo)

	playerSum := blackjack.AddCards(playerCardOne, playerCardTwo)
	fmt.Printf("Player's sum: %d.\n\n", playerSum)
	*/
}
