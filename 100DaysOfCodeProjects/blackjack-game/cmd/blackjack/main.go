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


	// Perhaps move this to blackjack.go file to test all cases appropriately?
	switch {
		case playerSum == 21:
			fmt.Printf("You win!\n")
			fmt.Printf("Player hand: %v. Player sum: %d.\n", playerHand, playerSum)
			fmt.Printf("Computer hand: %v. Dealer sum: %d.\n", dealerHand, dealerSum)
		case dealerSum == 21:
			fmt.Printf("You lose!\n")
			fmt.Printf("Player hand: %v. Player sum: %d.\n", playerHand, playerSum)
			fmt.Printf("Computer hand: %v. Dealer sum: %d.\n", dealerHand, dealerSum)
		case playerSum > dealerSum && playerSum <= 21 || dealerSum > 21:
			fmt.Printf("You win!\n")
			fmt.Printf("Player hand: %v. Player sum: %d.\n", playerHand, playerSum)
			fmt.Printf("Computer hand: %v. Dealer sum: %d.\n", dealerHand, dealerSum)
		case dealerSum > playerSum && dealerSum <= 21 || playerSum > 21:
			fmt.Printf("You lose!\n")
			fmt.Printf("Player hand: %v. Player sum: %d.\n", playerHand, playerSum)
			fmt.Printf("Computer hand: %v. Dealer sum: %d.\n", dealerHand, dealerSum)
		}

}
