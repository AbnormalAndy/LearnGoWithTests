package main

import (
	"averageheight"
	"fmt"
)

func main() {
	ExampleOne := []int{156, 178, 165, 171, 187}

	ExampleOneTotalHeight := float64(averageheight.CalculateTotalHeight(ExampleOne))
	ExampleOneTotalStudents := float64(averageheight.CalculateNumberOfStudents(ExampleOne))
	ExampleOneAverageHeight, _  := averageheight.CalculateAverage(ExampleOneTotalHeight, ExampleOneTotalStudents)
	fmt.Printf("\nThe average height for Example One is %d.\n", int(ExampleOneAverageHeight))

	ExampleTwo := []int{151, 145, 179}
	ExampleTwoTotalHeight := float64(averageheight.CalculateTotalHeight(ExampleTwo))
	ExampleTwoTotalStudents := float64(averageheight.CalculateNumberOfStudents(ExampleTwo))
	ExampleTwoAverageHeight, _ := averageheight.CalculateAverage(ExampleTwoTotalHeight, ExampleTwoTotalStudents)
	fmt.Printf("\nThe average height for Example Two is %d.\n\n", int(ExampleTwoAverageHeight))
}
	
