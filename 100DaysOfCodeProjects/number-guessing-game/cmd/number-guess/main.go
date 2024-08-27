package main

import (
	"bufio"
	"fmt"
	"numberguessinggame"
	"os"
	"strings"
)

func main() {
	var difficulty string

	fmt.Println("\nWelcome to the Number Guessing Game!\n")
	fmt.Println("I am thinking of a number between 1 and 100.\n")
		
	numberOfGuesses := 0
	finishDifficultyInput := false

	for finishDifficultyInput != true {
		
		fmt.Println("Choose a difficulty. Type 'easy' or 'hard':\n")
		
		inputDifficulty := bufio.NewReader(os.Stdin)
		difficulty, _ = inputDifficulty.ReadString('\n')
		difficulty = strings.TrimSuffix(difficulty, "\n")
		difficulty = strings.ToLower(difficulty)
		
		switch {
		case difficulty == "easy" || difficulty == "e":
			fmt.Println("\nYou have 10 attempts to guess the number.\n")
			numberOfGuesses = 10
			finishDifficultyInput = true
		case difficulty == "hard" || difficulty == "h":
			fmt.Println("\nYou have 5 attempts to guess the number.\n")
			numberOfGuesses = 5
			finishDifficultyInput = true
		default:
			fmt.Println("\nYou did not choose 'easy' or 'hard'.\n")
		}
	}

	finishGame := false

	for finishGame != true {
	
		// Use numberOfGuesses of 0 to break and try again.
		fmt.Println(difficulty)
		fmt.Println(numberOfGuesses)

		// Add a for loop to keep guessing.
		randomNumber := 10
		playerNumber := 5
		fmt.Println(numberguessinggame.EvaluateIfLower(randomNumber, playerNumber))
		finishGame = true
	}
}

