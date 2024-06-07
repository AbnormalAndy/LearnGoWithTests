package funwithfunctions

func double(x float64) float64 {
	return x * 2
}

func hello() {
	fmt.Println("Hi, I'm the function body!")
}

answer := double(2.5)

answer := 3 * (double(2.5) + 7)

lat, long, err := location()

// The defer Keyword
f, err := os.Open("testdata/somefile.txt")
if err != nil {
	return err
}
defer f.Close()
// Do stuff with f.

