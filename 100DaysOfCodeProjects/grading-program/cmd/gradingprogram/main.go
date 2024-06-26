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

	freddie := gradingprogram.Student{
		Name:	"Freddie",
		Score:	13,
		Grade:	"",
	}

	perfect := gradingprogram.Student{
		Name:	"Perfect",
		Score:	100,
		Grade:	"",
	}

	zero := gradingprogram.Student{
		Name:	"Zero",
		Score:	0,
		Grade:	"",
	}

	fmt.Println(harry)
	fmt.Println(ron)
	fmt.Println(hermione)
	fmt.Println(draco)
	fmt.Println(neville)
	fmt.Println(freddie)
	fmt.Println(perfect)
	fmt.Println(zero)

	harrygrade := gradingprogram.ScoreToGrade(harry)
	rongrade := gradingprogram.ScoreToGrade(ron)
	hermionegrade := gradingprogram.ScoreToGrade(hermione)
	dracograde := gradingprogram.ScoreToGrade(draco)
	nevillegrade := gradingprogram.ScoreToGrade(neville)
	freddiegrade := gradingprogram.ScoreToGrade(freddie)
	perfectgrade := gradingprogram.ScoreToGrade(perfect)
	zerograde := gradingprogram.ScoreToGrade(zero)

	fmt.Println(harrygrade)
	fmt.Println(rongrade)
	fmt.Println(hermionegrade)
	fmt.Println(dracograde)
	fmt.Println(nevillegrade)
	fmt.Println(freddiegrade)
	fmt.Println(perfectgrade)
	fmt.Println(zerograde)
}
