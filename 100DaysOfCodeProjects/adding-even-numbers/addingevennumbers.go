package addingevennumbers

func AddEvenNumbers(number int) int {
	sum := 0
	for i := range (number + 1) {
		if i % 2 == 0 {
			sum += i
		}
	}
	return sum
}
