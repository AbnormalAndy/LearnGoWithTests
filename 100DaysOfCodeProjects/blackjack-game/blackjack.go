package blackjack

import (
	"math/rand"
)

// Adds cards in slice and, if the sum is more than 21, will try convert an Ace to a 1.
// TO-DO: Figure out a way to evaluate greater than 21 hands.
// TO-DO: Convert 11 to 1 if hand greater than 21.
// TO-DO: Figure out a way to only conver 11 to 1 for only one of the 11's.
// Last test is failing.
func AddCards(hand []int) int {
	sum := 0
	for i := range hand {
		sum += hand[i]
	}

	if sum > 21 {
		sum = 0
		for i := range hand {
			if hand[i] == 11 {
				hand[i] = 1
			}
			sum += hand[i]
		}

		return sum
	}

	return sum
}

// Intakes a deck size of 13 and outputs a random card from that deck.
func DealCard(deckOne [13]int) int {
	randomCard := rand.Intn(len(deckOne))
	return deckOne[randomCard]
}
