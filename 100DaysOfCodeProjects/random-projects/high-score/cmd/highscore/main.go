package main

import (
	"fmt"
	"highscore"
)

func main() {
	listone := []int{57, 68, 34, 120, 23, 87, 101}
	listtwo := []int{21, 93, 77, 95, 52, 45, 79}

	highscoreone := highscore.FindHighScore(listone)
	highscoretwo := highscore.FindHighScore(listtwo)

	fmt.Printf("\nHigh Score Calculator\n\n")

	fmt.Printf("The high score for list one: %d.\n\n", highscoreone)
	fmt.Printf("The high score for list two: %d.\n\n", highscoretwo)
}
