package blackjack

import (
	"math/rand"
)

// Adds cards in slice and, if the sum is more than 21, will convert an Ace to a 1.
func AddCards(hand []int) int {
	sum := 0

	for i := range hand {
		sum += hand[i]
		if hand[i] == 11 {
		}
	}

	if sum > 21 {
		for i := range hand {
			if hand [i] == 11 && sum > 21 {
				sum -= 10
			}
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
