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
	fmt.Println("Choose a difficulty. Type 'easy' or 'hard':\n")
	
	inputDifficulty := bufio.NewReader(os.Stdin)
	difficulty, _ = inputDifficulty.ReadString('\n')
	difficulty = strings.TrimSuffix(difficulty, "\n")
	difficulty = strings.ToLower(difficulty)
	
	// May have to make this a switch.false
	if difficulty == "easy" || difficulty == "e" {
		fmt.Println("\nYou have 10 attempts to guess the number.\n")
	}
	
	if difficulty == "hard" || difficulty == "h" {
		fmt.Println("\nYou have 5 attempts to guess the number.\n")
	}

	fmt.Println(difficulty)

	// Add a for loop to keep guessing.
	randomNumber := 10
	playerNumber := 5
	fmt.Println(numberguessinggame.EvaluateIfLower(randomNumber, playerNumber))
}

