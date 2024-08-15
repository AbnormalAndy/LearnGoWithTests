package numberguessinggame

func EvaluateIfHigher(randomNumber, playerGuess int) bool {
	if randomNumber < playerGuess {
		return true
	}
	return false
}

func EvaluateIfLower(randomNumber, playerGuess int) bool {
	if randomNumber > playerGuess {
		return true
	}
	return false
}

func EvaluateIfEqual(randomNumber, playerGuess int) bool {
	if randomNumber == playerGuess {
		return true
	}
	return false
}
