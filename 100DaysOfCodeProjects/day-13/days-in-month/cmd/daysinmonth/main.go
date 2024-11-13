package main

import (
	"daysinmonth"
	"fmt"
)

func main() {
	fmt.Printf("Number of days in February 2020: %d.\n", daysinmonth.DaysInMonth(2020, 2))
	fmt.Printf("Number of days in February 2021: %d.\n", daysinmonth.DaysInMonth(2021, 2))
	fmt.Printf("Number of days in February 2022: %d.\n", daysinmonth.DaysInMonth(2022, 2))
	fmt.Printf("Number of days in February 2023: %d.\n", daysinmonth.DaysInMonth(2023, 2))
	fmt.Printf("Number of days in February 2024: %d.\n", daysinmonth.DaysInMonth(2024, 2))
}
