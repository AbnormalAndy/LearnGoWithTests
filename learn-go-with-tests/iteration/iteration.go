package iteration


const repeatedCount = 5

func Repeat(letter string) string {
	repeated := ""
	for i := 0; i < repeatedCount; i++ {
		repeated += letter
	}
	return repeated
}
