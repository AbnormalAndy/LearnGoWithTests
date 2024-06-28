package main

import (
	"fmt"
	"gradingprogram"
)

func main() {
	students := []gradingprogram.Student{
		{
			Name:	"Harry",
			Score:	81,
			Grade:	"",
		},
		{
			Name:	"Ron",
			Score:	78,
			Grade: "",
		},
		{
			Name:	"Hermione",
			Score:	99,
			Grade:	"",
		},
		{
			Name:	"Draco",
			Score:	74,
			Grade:	"",
		},
		{
			Name:	"Neville",
			Score:	62,
			Grade:	"",
		},
		{
			Name:	"Freddie",
			Score:	13,
			Grade:	"",
		},
		{
			Name:	"Perfect",
			Score:	100,
			Grade:	"",
		},
		{
			Name:	"Zero",
			Score:	0,
			Grade:	"",
		},
		{
			Name:	"Negative",
			Score:	-1,
			Grade:	"",
		},
	}

	// Perhaps put results in another array?
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
		fmt.Println(gradingprogram.ScoreToGrade(students[i]))
	}

	/*
	fmt.Println(harry)
	fmt.Println(ron)
	fmt.Println(hermione)
	fmt.Println(draco)
	fmt.Println(neville)
	fmt.Println(freddie)
	fmt.Println(perfect)
	fmt.Println(zero)
	fmt.Println(negative)

	harrygrade := gradingprogram.ScoreToGrade(harry)
	rongrade := gradingprogram.ScoreToGrade(ron)
	hermionegrade := gradingprogram.ScoreToGrade(hermione)
	dracograde := gradingprogram.ScoreToGrade(draco)
	nevillegrade := gradingprogram.ScoreToGrade(neville)
	freddiegrade := gradingprogram.ScoreToGrade(freddie)
	perfectgrade := gradingprogram.ScoreToGrade(perfect)
	zerograde := gradingprogram.ScoreToGrade(zero)
	negativegrade := gradingprogram.ScoreToGrade(negative)

	fmt.Println(harrygrade)
	fmt.Println(rongrade)
	fmt.Println(hermionegrade)
	fmt.Println(dracograde)
	fmt.Println(nevillegrade)
	fmt.Println(freddiegrade)
	fmt.Println(perfectgrade)
	fmt.Println(zerograde)
	fmt.Println(negativegrade)
	*/
}
