package fizzbuzz

// This is just taking a single number. Will filter for specific number in main.go.
// This will determine if that specific number should be Fizz Buzz or FizzBuzz.
func FizzBuzz(number int) string {
	switch {
	case number % 3 == 0 && number % 5 == 0:
		return "FizzBuzz"
	case number % 3 == 0:
		return "Fizz"
	case number % 5 == 0:
		return "Buzz"
	}
	return "" // Would this be where I return an error? Should I error handle here? Will think about.
	// Will need to read more about switch statements.
}
