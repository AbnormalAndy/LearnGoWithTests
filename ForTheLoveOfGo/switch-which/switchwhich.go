package switchwhich

switch {
case x < 0 :
	fmt.Println("Negative")
case x > 0:
	fmt.Println("Positive")
default:
	fmt.Println("Zero")
}

switch x {
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")
case 3:
	fmt.Println("three")
}

switch x {
case 1, 2, 3:
	fmt.Println("One, Two, or Three")
}

switch x {
case 1:
	if SomethingWentWrong() {
		break
	}
	... // Otherwise carry on.
}


// The for Keyword
for x < 10 {
	fmt.Println("X is less than 10.")
	// Update X, perhaps.
}


// Using range to loop over collections.
for range employees {
	fmt.Println("Found another employee!")
}

