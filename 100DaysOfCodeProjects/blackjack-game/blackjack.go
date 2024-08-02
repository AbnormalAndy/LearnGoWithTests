package blackjack

import (
	"math/rand"
)

// Adds two cards together but, if the sum is more than 21, will try convert an Ace to a 1.
func AddCards(cardOne, cardTwo int) int {
	sum := cardOne + cardTwo

	if sum > 21 && cardOne == 11 {
		return 1 + cardTwo
	}

	if sum > 21 && cardTwo == 11 {
		return cardOne + 1
	}

	return sum
}

// Intakes a deck size of 13 and outputs a random card from that deck.
func DealCard(deckOne [13]int) int {
	randomCard := rand.Intn(len(deckOne))
	return deckOne[randomCard]
}
