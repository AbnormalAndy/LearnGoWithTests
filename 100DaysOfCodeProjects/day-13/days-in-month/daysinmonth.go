package daysinmonth

//import "fmt"

//var DaysInMonths = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func IsLeap(year int) bool {
	if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
		return true
	} else {
		return false
	}
}

func DaysInMonth(year int, month int) int {
	var DaysInMonths = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if IsLeap(year) == true {
		DaysInMonths[1] += 1
	}
	numberOfDays := DaysInMonths[month - 1]
	return numberOfDays
}
