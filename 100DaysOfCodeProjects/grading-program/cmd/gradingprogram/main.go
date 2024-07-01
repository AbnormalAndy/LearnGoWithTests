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
			Name:	"Mikko",
			Score:	90,
			Grade:	"",
		},
		{
			Name:	"Elvira",
			Score:	100,
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

	fmt.Println(students)

	for i := range students {
		students[i] = gradingprogram.ScoreToGrade(students[i]) 
	}

	fmt.Println(students)
}
