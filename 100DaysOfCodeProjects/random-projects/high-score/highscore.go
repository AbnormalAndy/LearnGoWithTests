package highscore

func FindHighScore(numbers []int) int {
	var highscore int
	for i := range len(numbers) {
		if highscore < numbers[i] {
			highscore = numbers[i]
		}
	}
	return highscore
}
