package main

import (
	"fmt"
	"gradingprogram"
)

func main() {
	hermione := gradingprogram.Student{
		Name:	"Hermione",
		Score:	99,
		Grade:	"",
	}

	fmt.Println(hermione)

	hermionegrade := gradingprogram.ScoreToGrade(hermione)

	fmt.Println(hermionegrade)
}
