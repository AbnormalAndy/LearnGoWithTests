package main

import (
	"fmt"
	"wordguessinggame"
)


var words = []string{
	"apple",
	"banana",
	"broccoli",
	"cherry",
	"fig",
	"garlic",
	"kiwi",
	"kale",
	"lemon",
	"leek",
	"mango",
	"yams",
}


func main() {
	fmt.Print(words)
	fmt.Print(wordguessinggame.MakePlayerWord(words[0]))
}
