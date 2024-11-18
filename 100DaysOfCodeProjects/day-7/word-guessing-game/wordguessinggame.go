package wordguessinggame


func MakePlayerWord(word string) []string {
	player_word := []string{}
	for range(len(word)) {
		player_word = append(player_word, "_")
	}
	return player_word
}

