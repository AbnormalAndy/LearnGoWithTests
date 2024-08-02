package blackjack

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

