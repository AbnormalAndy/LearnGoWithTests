package controlflow

// Assigning More Than One Value
a, b, c := 1, 2, 3


// The Blank Identifier
_ = "Hello"
fmt.Println(_)
// Compile error: cannot use _ as value.

price, ok := menu["eggs"]
// Compile error: price declared but not used.

_, ok := menu["eggs"]



// Increment and Decrement Statements
x++

x := y++
// Syntax error: unexpected ++ at end of statement.

x--


// if Statements
if x > 0 {
	fmt.Println("X is positive.")
}


// The Happy Path
if x <= 0 {
	fmt.Println("Nope, x is zero or negative.")
	return false
}
if x % 2 != 0 {
	fmt.Println("Positive, but odd. Too bad.")
}
fmt.Println("Yay! That's a valid input.")
return true


// else Branches
fmt.Println("Let's see what the sign of x is:")
if x <= 0 {
	fmt.Println("X is zero or negative.")
} else {
	fmt.Println("X is positive.")
}
fmt.Println("Well, that clears that up!")


// Early return
if x <= 0 {
	fmt.Println("Nope, x is zero or negative.")
	return false
}


// And, or, and not
if x > 0 && x % 2 == 0 {
	fmt.Println("Positive and Even")
}

if x > 0 || x % 2 == 0 {
	fmt.Println("Positive or Even (or Both)")
}


// bool Variables
_, ok := menu["eggs"]
if ok {
	fmt.Println("Eggs are on the menu!")
}


// Compound if Statements
if _, ok := menu["eggs"]; ok {
	fmt.Println("Eggs are on the menu!")
}

//if STATEMENT; CONDITION {
//	...
//}

if _, ok := menu["eggs"]; ok {
	fmt.Println("Eggs are on the menu!")
}
fmt.Println(ok)
// error: undefined: ok

if err := doStuff(); err != nil {
	fmt.Println("Oh no.")
}

if err := apply(func(x int) error {
	if x <= 0 {
		return errors.New("Prognosis negative.")
	}
	return nil
}, -1); err != nil {
	fmt.Println("Some error happened.")
}


