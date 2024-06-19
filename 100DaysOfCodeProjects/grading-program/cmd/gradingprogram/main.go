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

	draco := gradingprogram.Student{
		Name:	"Draco",
		Score:	74,
		Grade:	"",
	}

	fmt.Println(hermione)
	fmt.Println(draco)

	hermionegrade := gradingprogram.ScoreToGrade(hermione)
	dracograde := gradingprogram.ScoreToGrade(draco)

	fmt.Println(hermionegrade)
	fmt.Println(dracograde)
}
