package iteration


func Repeat(letter string, number int) string {
	repeated := ""
	for i := 0; i < number; i++ {
		repeated += letter
	}
	return repeated
}
