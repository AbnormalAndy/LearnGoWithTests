package main

import (
	"blackjack"
	"fmt"
)

func main() {
	deck := [13]int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

	dealerCardOne := blackjack.DealCard(deck)
	dealerCardTwo := blackjack.DealCard(deck)
	fmt.Printf("Dealer's card one: %d; card two: %d.\n", dealerCardOne, dealerCardTwo)

	dealerSum := blackjack.AddCards(dealerCardOne, dealerCardTwo)
	fmt.Printf("Dealer's sum: %d.\n\n", dealerSum)

	if dealerSum < 17 {
		dealerCardThree := blackjack.DealCard(deck)
		fmt.Printf("Dealer's card three: %d.\n", dealerCardThree)
		dealerSum := blackjack.AddCards(dealerSum, dealerCardThree)
		fmt.Printf("Dealer's sum: %d.\n\n", dealerSum)
		// Drew an 11.
		// If not greater than player, convert to 1.
		// If greater than player with 11, wins.
	}

	playerCardOne := blackjack.DealCard(deck)
	playerCardTwo := blackjack.DealCard(deck)
	fmt.Printf("Player's card one: %d; card two: %d.\n", playerCardOne, playerCardTwo)

	playerSum := blackjack.AddCards(playerCardOne, playerCardTwo)
	fmt.Printf("Player's sum: %d.\n\n", playerSum)
}
