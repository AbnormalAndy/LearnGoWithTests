package main

import (
	"fmt"
	"gradingprogram"
)

func main() {
	ron := gradingprogram.Student{
		Name:	"Ron",
		Score:	78,
		Grade: "",
	}

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

	fmt.Println(ron)
	fmt.Println(hermione)
	fmt.Println(draco)

	rongrade := gradingprogram.ScoreToGrade(ron)
	hermionegrade := gradingprogram.ScoreToGrade(hermione)
	dracograde := gradingprogram.ScoreToGrade(draco)

	fmt.Println(rongrade)
	fmt.Println(hermionegrade)
	fmt.Println(dracograde)
}
