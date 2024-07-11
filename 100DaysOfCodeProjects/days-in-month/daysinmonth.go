package daysinmonth

var DaysInMonths = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func IsLeap(year int) bool {
	if year % 4 == 0 {
		if year % 100 == 0 {
			if year % 400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}

func DaysInMonth(year int, month int) int {
	if IsLeap(year) == true {
		DaysInMonths[1] += 1
	}
	numberOfDays := DaysInMonths[month - 1]
	return numberOfDays
}
