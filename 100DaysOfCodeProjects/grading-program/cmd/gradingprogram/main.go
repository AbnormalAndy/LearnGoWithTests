package main

import (
	"fmt"
	"gradingprogram"
)

func main() {
	harry := gradingprogram.Student{
		Name:	"Harry",
		Score:	81,
		Grade:	"",
	}

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

	neville := gradingprogram.Student{
		Name:	"Neville",
		Score:	62,
		Grade:	"",
	}

	fmt.Println(harry)
	fmt.Println(ron)
	fmt.Println(hermione)
	fmt.Println(draco)
	fmt.Println(neville)

	harrygrade := gradingprogram.ScoreToGrade(harry)
	rongrade := gradingprogram.ScoreToGrade(ron)
	hermionegrade := gradingprogram.ScoreToGrade(hermione)
	dracograde := gradingprogram.ScoreToGrade(draco)
	nevillegrade := gradingprogram.ScoreToGrade(neville)

	fmt.Println(harrygrade)
	fmt.Println(rongrade)
	fmt.Println(hermionegrade)
	fmt.Println(dracograde)
	fmt.Println(nevillegrade)
}
